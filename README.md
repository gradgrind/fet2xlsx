# fet2xlsx

> Convert [`FET`](https://lalescu.ro/liviu/fet/) (timetable) result files to simple Excel tables.

This takes a `FET` result file, the one containing the full data set, ending with "_data_and_timetable.fet", and produces two Excel documents, one for the teachers and one for the student groups. The results are placed in the same folder as the source file. It won't work on other types of `FET` result file.

As well as the individual sheets for each teacher or group, there is a sort of overview sheet. In the teachers document this lists, for each teacher, the student-groups in each hour. In the student-groups document this lists, for each student-group, the subject in each hour.

The students are divided into years (classes, grades, whatever ...) and the groups within each year. I use "year" and "group" as they are used in `FET`. The activities shown for a year are not repeated in the groups belonging to the year. `FET` subgroups are ignored here.

There are many ways in which the data could be presented. This is currently just a "proof of concept". It shouldn't be difficult to construct the tables differently.

## Using the program

### Command line (Linux)

First make the program executable, then

```
./fet2xlsx path/to/test1_data_and_timetable.fet
```

### GUI

Just run the GUI version – fet2xlsx.exe on Windows – and select the file to convert.

## Sketch of the inner workings

This program is written in Go and should run on Linux, Windows and MacOS.

The input file is XML, which is parsed to Go structures, in package "readfet". The GUI version uses the package ["zenity"](https://github.com/ncruces/zenity) to provide simple dialogs.

The generation of Excel files is done in package "fet2xlsx", using the ["excelize"](https://github.com/xuri/excelize/v2) package. Basically, the list of activities is read, filling the cells of the sheets with the available data. The activity times and the rooms are specified in constraints, which are read beforehand to extract the necessary information.
