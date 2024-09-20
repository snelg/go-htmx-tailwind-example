package main

import "strconv"

type Companies map[string]Company

type CompanyDB struct {
	companies Companies
	incr      int
}

var companyDB CompanyDB

type Company struct {
	ID      string
	Company string
	Contact string
	Country string
}

func init() {
	companies := Companies{
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

	companyDB = CompanyDB{
		companies: companies,
		incr:      len(companies),
	}
}

func (c *CompanyDB) getAll() Companies {
	return c.companies
}

func (c *CompanyDB) getById(id string) Company {
	return c.companies[id]
}

func (c *CompanyDB) update(id string, company Company) {
	c.companies[id] = company
}

func (c *CompanyDB) add(company Company) {
	c.incr++
	company.ID = strconv.Itoa(c.incr)
	c.companies[company.ID] = company
}

func (c *CompanyDB) delete(id string) {
	delete(c.companies, id)
}
