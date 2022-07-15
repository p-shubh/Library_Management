package main

// var books_csv = readCsvFile("./books.csv")

// func importcsv(books_csv [][]string) {
// 	// we have to take the data from 1 to 1000
// 	// using loop

// 	for i := 1; i < 4000; i++ {

// 		// insert sql statement
// 		sql_insert := `insert into books_detail(book_title, book_author,book_cover_image,book_id) values($1,$2,$3,$4)`

// 		_, err := DB.Exec(sql_insert, books_csv[i][1], books_csv[i][2], books_csv[i][7], books_csv[i][0])

// 		if err != nil {
// 			log.Fatal("error: ", err)
// 		}
// 	}
// }

// func readCsvFile(filePath string) [][]string {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal("Unable to read input file "+filePath, err)
// 	}
// 	defer f.Close()

// 	csvReader := csv.NewReader(f)
// 	csvReader.Comma = ';' //A comma-separated values (CSV) file is a delimited text file that uses a comma to separate values.
// 	csvReader.FieldsPerRecord = -1
// 	csvReader.LazyQuotes = true
// 	books_csv, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal("Unable to parse file as CSV for "+filePath, err)
// 	}

// fmt.Println(books_csv)

// 	return books_csv
// }
