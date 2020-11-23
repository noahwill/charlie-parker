package seeder

import "charlie-parker/pkg/types"

var seedData = []types.CreateRateInput{
	{
		Days:  "mon,tues,thurs",
		Times: "0900-2100",
		TZ:    "America/Chicago",
		Price: 1500,
	},
	{
		Days:  "fri,sat,sun",
		Times: "0900-2100",
		TZ:    "America/Chicago",
		Price: 2000,
	},
	{
		Days:  "wed",
		Times: "0600-1800",
		TZ:    "America/Chicago",
		Price: 1750,
	},
	{
		Days:  "mon,wed,sat",
		Times: "0100-0500",
		TZ:    "America/Chicago",
		Price: 1000,
	},
	{
		Days:  "sun,tues",
		Times: "0100-0700",
		TZ:    "America/Chicago",
		Price: 925,
	},
}
