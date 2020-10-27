package entity

type Result struct {
	Status      int32
	Value       interface{}
	Desc        string
	Exception   interface{}
	Attachments map[string]interface{}
}

type Invocation struct {

}