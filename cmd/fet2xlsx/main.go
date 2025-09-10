package main

import (
	"fmt"
	"strings"

	"fet2xlsx/fet2xlsx"
	"fet2xlsx/readfet"
	"flag"
	"path/filepath"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		if len(args) == 0 {
			panic("ERROR* No input file")
		}
		panic(fmt.Sprintf("*ERROR* Too many command-line arguments:\n  %+v\n", args))
	}
	abspath, err := filepath.Abs(args[0])
	if err != nil {
		panic(fmt.Sprintf("*ERROR* Couldn't resolve file path: %s\n", args[0]))
	}

	fetdata, err := readfet.ReadFet(abspath)
	if err != nil {
		panic(err)
	}

	// Generate output
	stempath := strings.TrimSuffix(abspath, filepath.Ext(abspath))
	fet2xlsx.GetTeachers(fetdata)
	fet2xlsx.GetStudentGroups(fetdata)
	activities := fet2xlsx.GetActivityData(fetdata)
	opath, err := fet2xlsx.TeachersActivities(fetdata, activities, stempath)
	if err == nil {
		fmt.Printf("Generated: %s\n", opath)
	} else {
		fmt.Println(err)
	}
	opath, err = fet2xlsx.StudentsActivities(fetdata, activities, stempath)
	if err == nil {
		fmt.Printf("Generated: %s\n", opath)
	} else {
		fmt.Println(err)
	}
}
