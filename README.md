# Stock Position APP

- Items inventory app to provide a basic stock movement support.  
- Learning about Clean Architecture/Hexagonal Architecture/Domain Driven Design 
- Golang practices  

## Technologies 
- Go v1.19.4
- MySQL driver v1.7.0

## Domain
- Stock Position: 
````json
{
    "id": "string",
	"item_id": "string",
	"facility_id": "uint",
	"on_hand_qty": "int",
	"unavl_qty": "int",
	"transaction_id": "string",
	"position_date": "string", // ISO 8601 - UTC: ${yyyy-MM-ddTHH:mm:ss-Z}
}
````    
#### **Use Cases** 

- **DoMovement:** Change stock position data using basic buckets: On hand Quantity and Unavailable Quantity

- **FindByItemIDAndFacilityID:** Get stock position data searching by Item ID and Facility ID





