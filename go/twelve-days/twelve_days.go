package twelve

type liryc struct {
	vers    int
	content string
}

var daysChristmas = []liryc{
	{
		vers:    1,
		content: "On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree.",
	},
	{
		vers:    2,
		content: "On the second day of Christmas my true love gave to me: two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    3,
		content: "On the third day of Christmas my true love gave to me: three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    4,
		content: "On the fourth day of Christmas my true love gave to me: four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    5,
		content: "On the fifth day of Christmas my true love gave to me: five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    6,
		content: "On the sixth day of Christmas my true love gave to me: six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    7,
		content: "On the seventh day of Christmas my true love gave to me: seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    8,
		content: "On the eighth day of Christmas my true love gave to me: eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    9,
		content: "On the ninth day of Christmas my true love gave to me: nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    10,
		content: "On the tenth day of Christmas my true love gave to me: ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    11,
		content: "On the eleventh day of Christmas my true love gave to me: eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
	{
		vers:    12,
		content: "On the twelfth day of Christmas my true love gave to me: twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	},
}

func Song() string {
	var ret string
	for idx, ve := range daysChristmas {
		if idx < len(daysChristmas)-1 {
			ret += ve.content + "\n"
		} else {
			ret += ve.content
		}
	}
	return ret
}

func Verse(vers int) string {
	vers--
	for idx, ve := range daysChristmas {
		if idx == vers {
			return ve.content
		}
	}
	return ""
}
