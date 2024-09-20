package main

import (
	"net/http"
)

// Delete -> DELETE /company/{id} -> delete, companys.html

// Edit   -> GET /company/edit/{id} -> row-edit.html
// Save   ->   PUT /company/{id} -> update, row.html
// Cancel ->	 GET /company/{id} -> nothing, row.html

// Add    -> GET /company/add/ -> companys-add.html (target body with row-add.html and row.html)
// Save   ->   POST /company -> add, companys.html (target body without row-add.html)
// Cancel ->	 GET /company -> nothing, companys.html

func index(_ *http.Request) *Response {
	return HTML("index.html", data)
}

func companyAdd(_ *http.Request) *Response {
	return HTML("company-add.html", data)
}

func companyEdit(r *http.Request) *Response {
	id := r.PathValue("id")
	row := getCompanyByID(id)
	return HTML("row-edit.html", row)
}

func companies(r *http.Request) *Response {
	return HTML("companies.html", data)
}

func companyGet(r *http.Request) *Response {
	id := r.PathValue("id")
	row := getCompanyByID(id)
	return HTML("row.html", row)
}

func companyPut(r *http.Request) *Response {
	id := r.PathValue("id")
	row := getCompanyByID(id)
	r.ParseForm()
	row.Company = r.Form.Get("company")
	row.Contact = r.Form.Get("contact")
	row.Country = r.Form.Get("country")
	updateCompany(id, row)
	return HTML("row.html", row)
}

func companyPost(r *http.Request) *Response {
	row := Company{}
	r.ParseForm()
	row.Company = r.Form.Get("company")
	row.Contact = r.Form.Get("contact")
	row.Country = r.Form.Get("country")
	addCompany(row)
	return HTML("companies.html", data)
}

func companyDelete(r *http.Request) *Response {
	id := r.PathValue("id")
	deleteCompany(id)
	return HTML("companies.html", data)
}
