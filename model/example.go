package model

import (
	"context"
	"template-api-gecho/config"

	//"template-api-gecho/entity"
	"fmt"
)

type Models struct {
	db config.Database
}

type Task struct {
	//	Id_user  int    `json:"id_user"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginTask struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CekTask struct {
	Id_user int `json:"id_user"`
	//Username string `json:"username"`
	Password string `json:"password"`
}
type UserTask struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

//registration
func (ExampleModel Models) Insert_users(Regis Task) bool {
	sqlStatement := "INSERT INTO users (email, username, password) " +
		"VALUES ($1, $2, $3)"
	res, err := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement,
		Regis.Email,
		Regis.Username,
		Regis.Password,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(res)
		return true
	}
}

func (ExampleModel Models) CekEmailUser(Email string) bool {
	// CEK EMAIL TERDAFTAR
	sqlStatement2 := "SELECT * FROM users " +
		"WHERE email=$1"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Email,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(res2)
	} else {
		fmt.Println(err2)
	}

	task := UserTask{}
	// for res2.Next() {
	// 	err2 := res2.Scan(&task.Email, &task.Username)
	// 	// Exit if we get an error
	// 	if err2 != nil {
	// 		fmt.Println(err2)
	// 	}
	// }

	if task.Email != "" {
		return false
	}
	return true
}

//login
func (ExampleModel Models) Ceklogin(Email string, Password string) bool {
	// CEK EMAIL TERDAFTAR
	sqlStatement2 := "SELECT  users.id_user, users.password FROM users " +
		"WHERE id_user=$1 AND password=$2 "
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Email,
		Password,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(res2)
	} else {
		fmt.Println(err2)
	}

	task := CekTask{}
	for res2.Next() {
		err2 := res2.Scan(&task.Id_user, &task.Password)
		// Exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	if task.Id_user != task.Id_user {
		return false
	} else if task.Password != Password {
		return false
	}
	return true
}

// Get Example Post
// func (ExampleModel ExampleModel) GetPosts(secret string) []entity.ExampleEntity {
// 	posts := []entity.ExampleEntity{
// 		{
// 			Title:       "NewsOne",
// 			Description: "NewsOneDescription",
// 		},
// 		{
// 			Title:       "NewsTwo",
// 			Description: "NewsTwoDescription",
// 		},
// 		{
// 			Title:       secret,
// 			Description: "NewsTwoDescription",
// 		},
// 	}

// 	return posts
// }

// Example For Insert Update and Delete
//Get Example Post
// func (UserModel UserModel) IsEmailOrNoTelponExist(identity string) (bool, error) {
// 	sqlStatement := "SELECT is_email_no_telp_exist($1)"
// 	isExist := false

// 	err := UserModel.db.GetDatabase().QueryRow(context.Background(),
// 		sqlStatement,
// 		identity,
// 	).Scan(&isExist)

// 	if err != nil {
// 		utils.LogError(err, utils.DetailFunction())
// 	}

// 	return isExist, err
// }

//Example For Get From DB
// func (UserModel UserModel) GetUserBussinessTypesByUserID(userID string) (entity.BussinessTypeList, error) {
// 	sqlStatement := "SELECT * FROM get_jenis_usaha($1)"

// 	bisnisTypeList := entity.BussinessTypeList{}
// 	res, err := UserModel.db.GetDatabase().Query(context.Background(),
// 		sqlStatement,
// 		userID,
// 	)

// 	defer utils.ResClose(res)

// 	if err != nil {
// 		utils.LogError(err, utils.DetailFunction())
// 		return bisnisTypeList, err
// 	}

// 	for res.Next() {
// 		tipeBisnis := entity.BussinessType{}
// 		err := res.Scan(
// 			&tipeBisnis.UserBussinessTypeID,
// 			&tipeBisnis.UserBussinessType,
// 		)
// 		// Exit if we get an error
// 		if err != nil {
// 			utils.LogError(err, utils.DetailFunction())
// 			return bisnisTypeList, err
// 		}
// 		bisnisTypeList.BussinessType = append(bisnisTypeList.BussinessType, tipeBisnis)
// 	}

// 	return bisnisTypeList, nil
// }

//mulai coment
type TaskProject struct {
	Id_project int    `json:"id_project"`
	Nama       string `json:"name"`
	Deskripsi  string `json:"deskripsi"`
	Id_user    int    `json:"id_user"`
	Aktor      string `json:"Aktorr"`
}

type TaskProjectView struct {
	Id_project int    `json:"id_project"`
	Aktor      string `json:"Aktorr"`
	Nama       string `json:"nama"`
	Deskripsi  string `json:"deskripsi"`
}

type ProjectsView struct {
	ProjectView []TaskProjectView `json:"project_view"`
}

//create project
func (ExampleModel Models) InsertTbl_Project(Create TaskProject) bool {

	sqlStatement2 := "INSERT INTO project (nama, description ) " +
		"VALUES ($1,$2)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Nama,
		Create.Deskripsi,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

func (ExampleModel Models) InsertTbl_Aktort(Create TaskProject) bool {

	sqlStatement2 := "INSERT INTO aktor (aktor ) " +
		"VALUES ($1)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Aktor,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//view project
func (ExampleModel Models) ViewProject(View TaskProjectView) ProjectsView {

	sqlStatement3 := "SELECT *  FROM project " +
		"WHERE project.id_project=$1 "
	res3, err3 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := ProjectsView{}

	for res3.Next() {
		task := TaskProjectView{}
		err3 := res3.Scan(&task.Nama, &task.Deskripsi, task.Aktor)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.ProjectView = append(result.ProjectView, task)
	}

	return result

}

type UseCasetask struct {
	Cases          string `json:"cases"`
	Tahapan        string `json:"tahapan"`
	Base_url       string `json:"base_url"`
	Pattern        string `json:"pattren"`
	Type           string `json:"type"`
	Function_query string `json:"function_query"`
	Base_query     string `json:"base_query"`
	Tipe_data      string `json:"tipe_data"`
	Request_param  string `json:"request_param"`
	Request_body   string `json:"request_body"`
	Header         string `json:"header"`
	Res            string `json:"res"`
	Data_object    string `json:"data_object"`
}

//create form usecase
func (ExampleModel Models) InsertTbl_cases(Create UseCasetask) bool {

	sqlStatement2 := "INSERT INTO cases (cases, tahapan  ) " +
		"VALUES ($1,$2)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Cases,
		Create.Tahapan,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//create form dfs
func (ExampleModel Models) InsertTbl_dfs(Create UseCasetask) bool {

	sqlStatement2 := "INSERT INTO dfs (base_url,pattern,type,function_query,base_query ) " +
		"VALUES ($1,$2,$3,$4,$5)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Base_url,
		Create.Pattern,
		Create.Type,
		Create.Function_query,
		Create.Base_query,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//create form Request param
func (ExampleModel Models) InsertTbl_RP(Create UseCasetask) bool {

	sqlStatement2 := "INSERT INTO request_param (request_param, tipe_data ) " +
		"VALUES ($1,$2)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Request_param,
		Create.Tipe_data,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//create form Header
func (ExampleModel Models) InsertTbl_Header(Create UseCasetask) bool {

	sqlStatement2 := "INSERT INTO header (header,tipe_data) " +
		"VALUES ($1,$2)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Header,
		Create.Tipe_data,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//create form Res
func (ExampleModel Models) InsertTbl_Res(Create UseCasetask) bool {

	sqlStatement2 := "INSERT INTO res (res,tipe_data ) " +
		"VALUES ($1,$2)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Res,
		Create.Tipe_data,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//create form data object
func (ExampleModel Models) InsertTbl_DO(Create UseCasetask) bool {

	sqlStatement2 := "INSERT INTO data_object (data_object, tipe_data ) " +
		"VALUES ($1,$2)"
	res2, err2 := ExampleModel.db.GetDatabase().Query(context.Background(), sqlStatement2,
		Create.Data_object,
		Create.Tipe_data,
	)
	defer ExampleModel.db.GetDatabase().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}
