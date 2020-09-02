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
		col := sheet.Cell(0, i).Value
		if col != "" {
			columns = append(columns, col)
		} else {
			break
		}
	}

	return columns
}

func buildRecordMap(columns []string, row *xlsx.Row) map[string]string {
	record := make(map[string]string)

	if len(row.Cells) < len(columns) {
		logrus.WithFields(logrus.Fields{
			"row": row.Cells,
			"cells": len(row.Cells),
			"columns": len(columns),
		}).Error("Row has fewer cells than columns.")

		return nil
	}

	for i, col := range columns {
		cell := row.Cells[i].Value
		if cell != "" {
			record[col] = row.Cells[i].Value
		}
	}

	return record
}

func buildTableMap(sheet *xlsx.Sheet) []map[string]string {
	cols := getColumnNames(sheet)

	if len(cols) == 0 {
		logrus.WithFields(logrus.Fields{
			"sheet": sheet.Name,
		}).Error("Sheet has no defined column names.")

		return nil
	}

	var table []map[string]string
	logrus.WithFields(logrus.Fields{
		"columns": len(cols),
		"names": cols,
	}).Debug("Defined table columns retrieved.")

	for i := 1; i < sheet.MaxRow; i++ {
		if len(sheet.Row(i).Cells) == 0 {
			continue
		}
		logrus.WithFields(logrus.Fields{
			"row": sheet.Row(i).Cells,
		}).Debug("Processing row")
		record := buildRecordMap(cols, sheet.Row(i))
		if len(record) != 0 {
			table = append(table, record)
		}
	}

	return table
}

func buildSeedMap(sheet *xlsx.File) []TableMap {
	var seed []TableMap

	for _, s := range sheet.Sheets {
		logrus.WithFields(logrus.Fields{
			"sheet": s.Name,
		}).Info("Processing sheet")
		objects := buildTableMap(s)
		if objects != nil {
			seed = append(seed, TableMap{
				s.Name,
				objects,
			})
		}
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
			"workbook":    filename,
			"error":       err,
		}).Fatal("Failed to open Excel workbook.")
	}

	seed := buildSeedMap(workbook)
	if seed == nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"workbook": filename,
		}).Fatal("Workbook does not have any valid tables.")
	}

	if err := writeSeedFile(seed, env); err != nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"workbook":    filename,
			"error":       err,
		}).Fatal("Failed to write seed file.")
	} else {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"workbook":    filename,
			"output_file": fmt.Sprintf("%s/%s.yml", viper.GetString("SEED_DIRECTORY"), env),
		}).Info("Seed file written.")
	}
}
