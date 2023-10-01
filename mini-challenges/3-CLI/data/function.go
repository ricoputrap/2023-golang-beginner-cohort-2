package data

func GetPersonByID(id int, people []TPerson) *TPerson {
	for _, person := range people {
		if (person.ID == id) {
			return &person;
		}
	}
	return nil;
}

func GetPersonByNama(nama string, people []TPerson) *TPerson {
	for _, person := range people {
		if (person.Nama == nama) {
			return &person;
		}
	}
	return nil;
}