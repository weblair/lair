package seeding

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tealeg/xlsx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type TableMap struct {
	TableName string
	Objects []map[string]string
}

func getColumnNames(sheet *xlsx.Sheet) []string {
	var columns []string

	for i := 0; i < sheet.MaxCol; i++ {
		columns = append(columns, sheet.Cell(0, i).Value)
	}

	return columns
}

func buildRecordMap(columns []string, row *xlsx.Row) map[string]string {
	record := make(map[string]string)

	for i, c := range columns {
		record[c] = row.Cells[i].Value
	}

	return record
}

func buildTableMap(sheet *xlsx.Sheet) []map[string]string {
	cols := getColumnNames(sheet)

	var table []map[string]string

	for i := 1; i < sheet.MaxRow; i++ {
		table = append(table, buildRecordMap(cols, sheet.Row(i)))
	}

	return table
}

func buildSeedMap(sheet *xlsx.File) []TableMap {
	var seed []TableMap

	for _, s := range sheet.Sheets {
		logrus.WithFields(logrus.Fields{
			"sheet": s.Name,
		}).Info("Processing sheet")
		seed = append(seed, TableMap{
			s.Name,
			buildTableMap(s),
		})
	}

	return seed
}

// writeSeedFile iterates over the TableMap slice and outputs it to the YAML file.
// The reason we aren't marhsalling everything as one giant map into YAML is that the order of the sheets matter.
// Marshalling everything to a giant YAML object does not preserve the order the sheets are processed.
func writeSeedFile(seed []TableMap, env string) error {
	var out []byte

	for _, s := range seed {
		b, err := yaml.Marshal(s.Objects)
		if err != nil {
			return err
		}

		out = append(out, []byte(fmt.Sprintf("%s:\n", s.TableName))...)
		out = append(out, b...)
	}

	if err := ioutil.WriteFile(
		fmt.Sprintf("%s/%s.yml", viper.GetString("SEED_DIRECTORY"), env),
		out,
		0644,
	); err != nil {
		return err
	}

	return nil
}

func GenerateSeedFile(env string, filename string) {
	workbook, err := xlsx.OpenFile(filename)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"filename":    filename,
			"error":       err,
		}).Fatal("Failed to open Excel workbook.")
	}

	seed := buildSeedMap(workbook)

	if err := writeSeedFile(seed, env); err != nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"filename":    filename,
			"error":       err,
		}).Fatal("Failed to write seed file.")
	}
}
