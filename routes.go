package main

import (
	"net/http"
)

// Delete -> DELETE /company/{id} -> delete, companies.html

// Edit   -> GET /company/edit/{id} -> row-edit.html
// Save   ->   PUT /company/{id} -> update, row.html
// Cancel ->	 GET /company/{id} -> nothing, row.html

// Add    -> GET /company/add/ -> company-add.html (target body with row-add.html and row.html)
// Save   ->   POST /company -> add, companies.html (target body without row-add.html)
// Cancel ->	 GET /company -> nothing, companies.html

func index(_ *http.Request) *Response {
	return HTML("index.html", companyDB.getAll())
}

func companyAdd(_ *http.Request) *Response {
	return HTML("company-add.html", companyDB.getAll())
}

func companyEdit(r *http.Request) *Response {
	id := r.PathValue("id")
	row := companyDB.getById(id)
	return HTML("row-edit.html", row)
}

func companiesGet(r *http.Request) *Response {
	return HTML("companies.html", companyDB.getAll())
}

func companyGet(r *http.Request) *Response {
	id := r.PathValue("id")
	row := companyDB.getById(id)
	return HTML("row.html", row)
}

func companyPut(r *http.Request) *Response {
	id := r.PathValue("id")
	row := companyDB.getById(id)
	r.ParseForm()
	row.Company = r.Form.Get("company")
	row.Contact = r.Form.Get("contact")
	row.Country = r.Form.Get("country")
	companyDB.update(id, row)
	return HTML("row.html", row)
}

func companyPost(r *http.Request) *Response {
	row := Company{}
	r.ParseForm()
	row.Company = r.Form.Get("company")
	row.Contact = r.Form.Get("contact")
	row.Country = r.Form.Get("country")
	companyDB.add(row)
	return HTML("companies.html", companyDB.getAll())
}

func companyDelete(r *http.Request) *Response {
	id := r.PathValue("id")
	companyDB.delete(id)
	return HTML("companies.html", companyDB.getAll())
}
