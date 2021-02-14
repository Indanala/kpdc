package controller

import (
	"net/http"
	//"os"
	//"os"
	//"template-api-gecho/constant"
	"fmt"
	"template-api-gecho/model"
	"template-api-gecho/responsegraph"

	"github.com/labstack/echo"
)

type Controller struct {
	ma model.Models
}

//Task
type Task struct {
	Id_user  int    `json:"id_user"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginTask struct {
	Email string `json:"email"`
	//Username  string `json:"username"`
	Password string `json:"password"`
}
type RegisTask struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginGet struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//registration
func (Controller Controller) PostsRegis(c echo.Context) error {
	a := new(RegisTask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.Task{
		Email:    a.Email,
		Username: a.Username,
		Password: a.Password,
	}

	fmt.Println(a.Email)

	//cekEmail := Controller.ma.CekEmailUser(ab.Email)
	insertUser := Controller.ma.Insert_users(ab)
	//insertUserAuth := Controller.ma.InsertTbl_UserAuth(ab)

	//if cekEmail {

	if insertUser {
		getRegis := LoginGet{a.Email, a.Username, a.Password}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil Input Data",
			Data:    getRegis,
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal Input User",
		}
		return c.JSON(http.StatusOK, res)
	}

	// } else {

	// 	res := responsegraph.ResponseGenericIn{
	// 		Status:  "Error",
	// 		Message: "Email Telah Digunakan",
	// 	}
	// 	return c.JSON(http.StatusOK, res)

	// }
}

//login
func (Controller Controller) GetLogin(c echo.Context) error {
	id := new(LoginTask)
	if err := c.Bind(id); err != nil {
		return err
	}

	ab := model.Task{
		Email:    id.Email,
		Password: id.Password,
	}
	cekLogin := Controller.ma.Ceklogin(ab.Email, ab.Password)
	//positions := Controller.ma.GetPositionUserLogin(ab)
	//posts := Controller.ma.LoginTask(ab.Email, ab.Password)

	if cekLogin {
		getRegis := LoginTask{id.Email, id.Password}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil Input Data",
			Data:    getRegis,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Login Gagal ",
		}
		return c.JSON(http.StatusOK, res)
	}

}

// type ExampleController struct {
// 	model model.ExampleModel
// }

//Get Example Controller
// func (ExampleController ExampleController) GetPostsController(c echo.Context) error {
// 	secret := os.Getenv(constant.Authorization)

// 	posts := ExampleController.model.GetPosts(secret)
// 	res := responsegraph.ResponseGenericGet{
// 		Status:  constant.StatusSuccess,
// 		Message: "Berhasil Insert Data",
// 		Data:    posts,
// 	}

// 	return c.JSON(http.StatusOK, res)
// }

//mulai komen
type TaskProject struct {
	Id_project int    `json:"id_project"`
	Nama       string `json:"name"`
	Deskripsi  string `json:"deskripsi"`
	Id_user    int    `json:"id_user"`
	Aktor      string `json:"Aktorr"`
}
type CreateGets struct {
	Id_project int    `json:"id_project"`
	Nama       string `json:"name"`
	Deskripsi  string `json:"deskripsi"`
	Aktor      string `json:"Aktor"`
}
type ProjectViewget struct {
	Id_project int `json:"id_project"`
	Id_user    int `json:"id_user"`
}

//create project
func (Controller Controller) PostsCreate(c echo.Context) error {
	a := new(TaskProject)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.TaskProject{

		Nama:      a.Nama,
		Deskripsi: a.Deskripsi,
	}

	insertTbl_project := Controller.ma.InsertTbl_Project(ab)

	//if cekIdProject {

	if insertTbl_project {

		getCreate := CreateGets{Nama: a.Nama, Deskripsi: a.Deskripsi, Aktor: a.Aktor}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create project",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal create project",
		}
		return c.JSON(http.StatusOK, res)

	}

}

//view
func (Controller Controller) ViewProject(c echo.Context) error {
	a := new(ProjectViewget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.TaskProjectView{
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewProject(ab)

	res := responsegraph.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil tampilkan project",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)
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

type CreateGetform struct {
	Id_project int    `json:"id_project"`
	Cases      string `json:"cases"`
	Tahapan    string `json:"tahapan"`
	//Id_user     string `json:"id_user"`
}

//create form usecase
func (Controller Controller) PostsCreateForm(c echo.Context) error {
	a := new(UseCasetask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.UseCasetask{

		Cases:   a.Cases,
		Tahapan: a.Tahapan,
	}

	insertTbl_cases := Controller.ma.InsertTbl_cases(ab)

	if insertTbl_cases {

		getCreate := CreateGetform{Cases: a.Cases, Tahapan: a.Tahapan}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create project",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal create project",
		}
		return c.JSON(http.StatusOK, res)

	}

}

//create tabel dfs
func (Controller Controller) PostsCreatedfs(c echo.Context) error {
	a := new(UseCasetask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.UseCasetask{

		Base_url:       a.Base_url,
		Pattern:        a.Pattern,
		Type:           a.Type,
		Function_query: a.Function_query,
		Base_query:     a.Base_query,
	}

	insertTbl_dfs := Controller.ma.InsertTbl_dfs(ab)

	if insertTbl_dfs {

		getCreate := UseCasetask{Base_url: a.Base_url, Pattern: a.Pattern, Type: a.Type, Function_query: a.Function_query, Base_query: a.Base_query}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil isi data tabel dfs",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal Isi",
		}
		return c.JSON(http.StatusOK, res)

	}

}

//create Request param
func (Controller Controller) PostsCreateRequestParam(c echo.Context) error {
	a := new(UseCasetask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.UseCasetask{

		Request_param: a.Request_param,
		Tipe_data:     a.Tipe_data,
	}

	insertTbl_RP := Controller.ma.InsertTbl_RP(ab)

	if insertTbl_RP {

		getCreate := UseCasetask{Request_param: a.Request_param, Tipe_data: a.Tipe_data}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil isi data tabel RP",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal Isi",
		}
		return c.JSON(http.StatusOK, res)

	}

}

//create Header
func (Controller Controller) PostsCreateHeader(c echo.Context) error {
	a := new(UseCasetask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.UseCasetask{

		Header:    a.Header,
		Tipe_data: a.Tipe_data,
	}

	InsertTbl_Header := Controller.ma.InsertTbl_Header(ab)

	if InsertTbl_Header {

		getCreate := UseCasetask{Header: a.Header, Tipe_data: a.Tipe_data}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil isi data tabel Header",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal Isi",
		}
		return c.JSON(http.StatusOK, res)

	}

}

//create Res
func (Controller Controller) PostsCreateRes(c echo.Context) error {
	a := new(UseCasetask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.UseCasetask{

		Res:       a.Res,
		Tipe_data: a.Tipe_data,
	}

	InsertTbl_Res := Controller.ma.InsertTbl_Res(ab)

	if InsertTbl_Res {

		getCreate := UseCasetask{Res: a.Res, Tipe_data: a.Tipe_data}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil isi data tabel Res",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal Isi",
		}
		return c.JSON(http.StatusOK, res)

	}

}

//create Res
func (Controller Controller) PostsCreateDO(c echo.Context) error {
	a := new(UseCasetask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := model.UseCasetask{

		Data_object: a.Data_object,
		Tipe_data:   a.Tipe_data,
	}

	InsertTbl_DO := Controller.ma.InsertTbl_DO(ab)

	if InsertTbl_DO {

		getCreate := UseCasetask{Data_object: a.Data_object, Tipe_data: a.Tipe_data}
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil isi data tabel DO",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGenericIn{
			Status:  "Error",
			Message: "Gagal Isi",
		}
		return c.JSON(http.StatusOK, res)

	}

}
