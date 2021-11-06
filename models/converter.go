package models

type Converter struct {
	Number string `gorm:"primary_key;size:200" json:"number"`
}

func (converter *Converter) TableName() string {
	return "converter"
}

func (converter *Converter) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["number"] = converter.Number
	return resp

}
