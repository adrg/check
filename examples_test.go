package check_test

import (
	"fmt"
	"time"

	"github.com/adrg/check"
)

func ExampleRun() {
	name := "Bond, James Bond"
	email := "007@example.co.uk"
	hobbies := []string{"gadgets", "shaken drinks", "puns"}
	weaponCaliber := 7.65
	movies := 25
	today := time.Now()
	tomorrow := time.Now().AddDate(0, 0, 1)
	contacts := map[string]string{
		"M": "m@example.co.uk",
		"Q": "q@example.co.uk",
	}

	if err := check.Run(
		// Required checks.
		check.Required(name),
		check.Required(email),
		check.Required(movies),
		check.Required(hobbies),

		// Equality checks.
		check.Eq(name, "Bond, James Bond"),
		check.Eq(movies, 25),
		check.Eq(hobbies, []string{"gadgets", "shaken drinks", "puns"}),
		check.Ne(email, "m@example.co.uk"),
		check.Ne(contacts, map[string]string{"Dr. No": "no@example.com"}),
		check.Ne(tomorrow, today),

		// Numeric comparison checks.
		check.Lt(movies, 26),
		check.Lte(weaponCaliber, 7.65),
		check.Gt(tomorrow, today),
		check.Gte(movies, 25),
		check.Between(time.Now(), today, tomorrow),

		// Presence checks.
		check.In(movies, 20, 25, 30),
		check.In(hobbies[0], "martini", "gadgets", "cars"),
		check.NotIn(weaponCaliber, 418, 45, 99, 308),
		check.NotIn(hobbies, []string{"suits", "cards"}, []string{"losing"}),

		// Regular expression checks.
		check.Matches("Dr. No", `\w+\. \w+`, true),
		check.Matches("Nick Nack", `\D+\s{1}Nack`, true),

		// Common entity checks.
		check.IBAN("IE64IRCE92050112345678", true),
		check.VAT("NO939194428", true),
		check.Email(email, true),
		check.EmailList("M <m@example.co.uk>, Q <q@example.co.uk>", true),
		check.URL("https://bond.example.com", true),
		check.IP("127.0.0.1", true),
		check.MAC("A3:4D:7A:8A:50:B8", true),
	); err != nil {
		// Treat error.
		fmt.Println(err)
	}
}

func ExampleRequired() {
	var email string
	if err := check.Run(check.Required(email)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Output: empty argument
}

func ExampleEq() {
	if err := check.Run(check.Eq(3, 4)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Eq(1, 1),
		check.Eq("a", "a"),
		check.Eq([]string{"a", "b", "c"}, []string{"a", "b", "d"}),
		check.Eq([]string{}, nil),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `eq` comparison failed: `3` is not equal to `4`
	// `eq` comparison failed: `[a b c]` is not equal to `[a b d]`
}

func ExampleNe() {
	if err := check.Run(check.Ne(2, 2)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Ne(1, 2),
		check.Ne("a", "b"),
		check.Ne(map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 2: 2, 3: 4}),
		check.Ne([]string{"a", "b", "c"}, []string{"a", "b", "c"}),
		check.Ne([]string{}, []string{}),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `ne` comparison failed: `2` is equal to `2`
	// `ne` comparison failed: `[a b c]` is equal to `[a b c]`
}

func ExampleLt() {
	if err := check.Run(check.Lt(2, 2)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Lt(1, 2),
		check.Lt(6.7, 4.5),
		check.Lt(time.Now(), time.Now()),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `lt` comparison failed: `2` is not less than `2`
	// `lt` comparison failed: `6.7` is not less than `4.5`
}

func ExampleLte() {
	if err := check.Run(check.Lte(3, 2)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Lte(2, 2),
		check.Lte(4.6, 4.5),
		check.Lte(time.Now(), time.Now()),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `lte` comparison failed: `3` is not less than or equal to `2`
	// `lte` comparison failed: `4.6` is not less than or equal to `4.5`
}

func ExampleGt() {
	if err := check.Run(check.Gt(2, 2)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Gt(2, 1),
		check.Gt(3.1, 4.3),
		check.Gt(time.Now(), time.Now()),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `gt` comparison failed: `2` is not greater than `2`
	// `gt` comparison failed: `3.1` is not greater than `4.3`
}

func ExampleGte() {
	if err := check.Run(check.Gte(2, 3)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Gte(2, 2),
		check.Gte(3.1, 4.3),
		check.Gte(time.Now(), time.Now()),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `gte` comparison failed: `2` is not greater than or equal to `3`
	// `gte` comparison failed: `3.1` is not greater than or equal to `4.3`
}

func ExampleBetween() {
	if err := check.Run(check.Between(2, 3, 4)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Between(5, 1, 10),
		check.Between(2.3, 1.0, 5.0),
		check.Between(10.5, 11.2, 15.3),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `gte` comparison failed: `2` is not greater than or equal to `3`
	// `gte` comparison failed: `10.5` is not greater than or equal to `11.2`
}

func ExampleIn() {
	if err := check.Run(check.In("a", "b", "c", "d")); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.In(2, 2, 3, 4),
		check.In([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5}),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `in` comparison failed: `a` not in `[b c d]`
	// `in` comparison failed: `[1 2 3]` not in `[[2 3 4] [3 4 5]]`
}

func ExampleNotIn() {
	if err := check.Run(check.NotIn("a", "a", "c", "d")); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.NotIn(2, 1, 3, 4),
		check.NotIn([]int{1, 2, 3}, []int{2, 3, 4}, []int{1, 2, 3}),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `not in` comparison failed: `a` in `[a c d]`
	// `not in` comparison failed: `[1 2 3]` in `[[2 3 4] [1 2 3]]`
}

func ExampleMatches() {
	if err := check.Run(check.Matches("32", `\D+`, true)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Matches("abc", `\w+`, true),
		check.Matches("abc", `\d+`, true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// `32` does not match pattern `\D+`
	// `abc` does not match pattern `\d+`
}

func ExampleEmail() {
	if err := check.Run(check.Email("test.example.com", true)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.Email("", false),
		check.Email("Alice <aliceexample.com>", true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// invalid email address `test.example.com`
	// invalid email address `Alice <aliceexample.com>`
}

func ExampleEmailList() {
	if err := check.Run(
		check.EmailList("eve@example.com, Bob <bobexample.com>", true),
	); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.EmailList("Eve <eve@example.com>, Bob <bob@example.com>", true),
		check.EmailList("Alice <alice@example.com>", true),
		check.EmailList("", false),
		check.EmailList("Bob <bob@example.com>,,", true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// invalid email address `Bob<bobexample.com>`
	// invalid email address ``
}

func ExampleURL() {
	if err := check.Run(check.URL("test@example", true)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.URL("", false),
		check.URL("https://example com", true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// invalid URL `test@example`
	// invalid URL `https://example com`
}

func ExampleIBAN() {
	if err := check.Run(check.IBAN("ALB3520111", true)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.IBAN("SV43ACAT00000000000000123123", true),
		check.IBAN("", false),
		check.IBAN("00CY2100200195000035700123", true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// invalid IBAN `ALB3520111`
	// invalid IBAN `00CY2100200195000035700123`
}

func ExampleVAT() {
	if err := check.Run(check.VAT("ZY1234567", true)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.VAT("ATU00000024", true),
		check.VAT("", false),
		check.VAT("AT0000", true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// invalid VAT number `ZY1234567`
	// invalid VAT number `AT0000`
}

func ExampleIP() {
	if err := check.Run(check.IP("192.168.100.256", true)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.IP("127.0.0.1", true),
		check.IP("::1", true),
		check.IP("", false),
		check.IP("23.55.3212", true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// invalid IP address `192.168.100.256`
	// invalid IP address `23.55.3212`
}

func ExampleMAC() {
	if err := check.Run(check.MAC("00:0a:95:9d:68:16:00", true)); err != nil {
		// Treat error.
		fmt.Println(err)
	}

	// Run multiple checks.
	if err := check.Run(
		check.MAC("00:A0:C9:14:C8:29", true),
		check.MAC("5F-7C-F5-12-FF-E7", true),
		check.MAC("", false),
		check.MAC("77-6B-00--79-DF-4C", true),
	); err != nil {
		// Treat error
		fmt.Println(err)
	}

	// Output:
	// invalid mac address `00:0a:95:9d:68:16:00`
	// invalid mac address `77-6B-00--79-DF-4C`
}
