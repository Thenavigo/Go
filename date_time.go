package main

import(
	f "fmt"
	  "time"
)

var pl = f.Println

func main() {


	// ----- TIME -----
	// Get day, month, year and time data
	// Get current time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())


	// Set a location to get time
	loc, err := time.LoadLocation("Africa/Douala")
	if err != nil {
		pl(err)
	}
	f.Printf("Time in Douala %s\n", now.In(loc))

	// Change location to Shanghai
	loc, _ = time.LoadLocation("Asia/Shanghai")
	f.Printf("Time in Shanghai %s\n", now.In(loc))


	// Get times using different time standards
	locEST, _ := time.LoadLocation("EST")
	locUTC, _ := time.LoadLocation("UTC")
	locMST, _ := time.LoadLocation("MST")
	f.Printf("EST : %s\n", now.In(locEST))
	f.Printf("UTC : %s\n", now.In(locUTC))
	f.Printf("MST : %s\n", now.In(locMST))


	// Calculate time since birthdate
	// Year, month, day, hour, minute, second
	// nanosecond and time zone
	birthDate := time.Date(1974, time.December, 
		21, 11, 30, 10, 0, time.Local)

	// Get difference between past date and now
	diff := now.Sub(birthDate)

	// Difference in days
	f.Printf("Days Alive : %d days\n", 
		int(diff.Hours()/24))

	// Hours
	f.Printf("Hours Alive : %d hours\n", 
		int(diff.Hours()))
}