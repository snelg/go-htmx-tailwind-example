package main

import "strconv"

var data map[string]Company
var incr int

type Company struct {
	ID      string
	Company string
	Contact string
	Country string
}

func init() {
	data = map[string]Company{
		"1": {
			ID:      "1",
			Company: "Amazon",
			Contact: "Jeff Bezos",
			Country: "United States",
		},
		"2": {
			ID:      "2",
			Company: "Apple",
			Contact: "Tim Cook",
			Country: "United States",
		},
		"3": {
			ID:      "3",
			Company: "Microsoft",
			Contact: "Satya Nadella",
			Country: "United States",
		},
	}
	incr = len(data)
}

func getCompanyByID(id string) Company {
	return data[id]
}

func updateCompany(id string, company Company) {
	data[id] = company
}

func addCompany(company Company) {
	incr++
	company.ID = strconv.Itoa(incr)
	data[company.ID] = company
}

func deleteCompany(id string) {
	delete(data, id)
}
