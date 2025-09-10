package main

import (
	"fet2xlsx/fet2xlsx"
	"fet2xlsx/readfet"
	"strings"

	"github.com/ncruces/zenity"

	"path/filepath"
)

const defaultPath = ``

func main() {
	abspath, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{
				Name:     "FET result files",
				Patterns: []string{"*_data_and_timetable.fet"},
				CaseFold: false,
			},
		})
	if err != nil {
		panic(err)
	}

	fetdata, err := readfet.ReadFet(abspath)
	if err != nil {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon)
		return
	}

	// Generate output
	stempath := strings.TrimSuffix(abspath, filepath.Ext(abspath))
	fet2xlsx.GetTeachers(fetdata)
	fet2xlsx.GetStudentGroups(fetdata)
	activities := fet2xlsx.GetActivityData(fetdata)
	opath, err := fet2xlsx.TeachersActivities(fetdata, activities, stempath)
	if err == nil {
		zenity.Info("Generated: "+opath,
			zenity.Title("Information"),
			zenity.InfoIcon)
	} else {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon)
	}
	opath, err = fet2xlsx.StudentsActivities(fetdata, activities, stempath)
	if err == nil {
		zenity.Info("Generated: "+opath,
			zenity.Title("Information"),
			zenity.InfoIcon)
	} else {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon)
	}
}
