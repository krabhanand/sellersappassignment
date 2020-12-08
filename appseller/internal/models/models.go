package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Product struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Price       string `json:"price" bson:"price,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Reviews     string `json:"reviews" bson:"reviews,omitempty"`
	ImageURL    string `json:"imageurl" bson:"imageurl,omitempty"`
}

type ProductData struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	URL       string              `json:"url" bson:"url,omitempty"`
	Product   *Product            `json:"product" bson:"product,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
