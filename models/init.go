package models
	
 func InitData() []User{
	var data_users []User
	u0 := User{FirstName:"Mary", Email:"qwerty@gmail.com", Username: "MaryPoppins"} 
    	u1 := User{FirstName:"Jack", Email:"bride@gmail.com", Username: "JackTheRipper"} 
	u2 := User{FirstName:"Karl", Email:"notorious@gmail.com", Username: "Highness"} 
		
	data_users = append(data_users, u0)
	data_users = append(data_users, u1)
	data_users = append(data_users, u2)
	return data_users 
	}