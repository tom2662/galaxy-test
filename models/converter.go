package models

// Model
type Converter struct {
	Number int64 `gorm:"primary_key;auto_increment" json:"number"`
}

// TableName method that returns tablename of Question model
func (converter *Converter) TableName() string {
	return "converter"
}

//ResponseMap -> response map of Converter
func (converter *Converter) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["number"] = converter.Number
	return resp

}
