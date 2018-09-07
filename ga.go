package main

import (
	"fmt"
	"os"
	"log"
	"sort"
	"github.com/urfave/cli"
	"gopkg.in/ini.v1"
	"crypto/hmac"
	"crypto/sha1"
	"strings"
	"encoding/base32"
	"time"
	"github.com/atotto/clipboard"
	"strconv"
)


func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F

	// get a 32-bit (4-byte) chunk from the hash starting at offset
	hashParts := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// size to 6 digits
	// one million is the first number with 7 digits so the remainder
	// of the division will always return < 7 digits
	pwd := number % 1000000

	return pwd
}



func Cli() {

	app := cli.NewApp()
	app.Name = "Google Authentiator CLI"
	app.Usage = "Create, List, Delete, Copy your GA"
	app.Version = "0.0.1"
	app.Commands = []cli.Command {
		{
			Name: "list",
			Aliases: []string{"l"},
			Usage: "list | l",
			Description: "list all item",
			Category: "Show",
			Action: func(c *cli.Context) error {
				list()
				return nil
			},
		},
		{
			Name: "show",
			Aliases: []string{"s"},
			Usage: "show | s github.com",
			Description: "show all item",
			Category: "Show",
			Action: func(c *cli.Context) error {
				item := c.Args().First()
				show(item)
				return nil
			},
		},
		{
			Name: "copy",
			Aliases: []string{"c"},
			Usage: "copy | c github.com",
			UsageText: "copy ga to Clipboard",
			Category: "Show",
			Action: func(c *cli.Context) error {
				item := c.Args().First()
				copy(item)
				return nil
			},
		},
		{
			Name: "create",
			Aliases: []string{"r"},
			Usage: "create| r github.com xxxxxxxxxx",
			UsageText: "create a new ga Secrete item",
			Category: "Store",
			Action: func(c *cli.Context) error {
				item := c.Args().First()
				secret := c.Args()[1]
				create(item, secret)
				return nil
			},
		},
		{
			Name: "modify",
			Aliases: []string{"m"},
			Usage: "modify | m github.com xxxxxxxxx",
			UsageText: "Modify One item Secrete",
			Category: "Store",
			Action: func(c *cli.Context) error {
				item := c.Args().First()
				secret := c.Args()[1]
				modify(item, secret)
				return nil
			},
		},
		{
			Name: "delete",
			Aliases: []string{"d"},
			Usage: "delete | d github.com",
			UsageText: "Delete One item Secrete",
			Category: "Store",
			Action: func(c *cli.Context) error {
				delete(c.Args().First())
				return nil
			},
		},

	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("you can use list")
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	cli.HelpFlag = cli.BoolFlag {
		Name: "help, h",
		Usage: "Help!Help!",
	}

	cli.VersionFlag = cli.BoolFlag {
		Name: "print-version, v",
		Usage: "print version",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func create(item, secret string) error {
	file,er:=os.Open("/tmp/myga.ini")

	if er != nil {
		file, _ = os.Create("/tmp/myga.ini")
	}
	file.Close()

	cfg, err := ini.Load("/tmp/myga.ini")
	if err != nil {
		panic(err)
	}

	if _, err := cfg.GetSection(item); err == nil {
		fmt.Println(item + " is Exist!!!!")
		os.Exit(1)
	}

	cfg.NewSection(item)
	cfg.Section(item).Key("secret").SetValue(secret)
	cfg.SaveTo("/tmp/myga.ini")
	return nil
}

func show(item string) error {
	file,er:=os.Open("/tmp/myga.ini")

	if er!=nil && os.IsNotExist(er){
		fmt.Println("no data!!!")
		os.Exit(1)
	}
	file.Close()

	cfg, err := ini.Load("/tmp/myga.ini")
	if err != nil {
		panic(err)
	}


	section, err := cfg.GetSection(item)
	if err != nil {
		panic(err)
	}

	key, err := section.GetKey("secret")

	if err != nil {
		panic(err)
	}

	inputNoSpaces := strings.Replace(key.String(), " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	keyS, erro := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if erro != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// generate a one-time password using the time at 30-second intervals
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(keyS, toBytes(epochSeconds/30))

	secondsRemaining := 30 - (epochSeconds % 30)
	fmt.Printf("%06d (%d second(s) remaining)\n", pwd, secondsRemaining)
	return nil

}


func copy(item string) error {
	file,er:=os.Open("/tmp/myga.ini")

	if er!=nil && os.IsNotExist(er){
		fmt.Println("no data!!!")
		os.Exit(1)
	}
	file.Close()

	cfg, err := ini.Load("/tmp/myga.ini")
	if err != nil {
		panic(err)
	}


	section, err := cfg.GetSection(item)
	if err != nil {
		panic(err)
	}

	key, err := section.GetKey("secret")

	if err != nil {
		panic(err)
	}

	inputNoSpaces := strings.Replace(key.String(), " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	keyS, erro := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if erro != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// generate a one-time password using the time at 30-second intervals
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(keyS, toBytes(epochSeconds/30))

	secondsRemaining := 30 - (epochSeconds % 30)
	clipStr := strconv.Itoa(int(pwd))
	clipboard.WriteAll(clipStr)
	fmt.Printf("%06d (%d second(s) remaining)\n", pwd, secondsRemaining)
	return nil

}




func list() error {
	file,er:=os.Open("/tmp/myga.ini")

	if er!=nil && os.IsNotExist(er){
		fmt.Println("no data!!!")
		os.Exit(1)
	}
	file.Close()

	cfg, err := ini.Load("/tmp/myga.ini")
	if err != nil {
		panic(err)
	}


	sections := cfg.Sections()

	for _,v := range sections {
		fmt.Println(v.Name())
	}
	return nil

}




func delete(item string) error {
	file,er:=os.Open("/tmp/myga.ini")

	if er!=nil && os.IsNotExist(er){
		file, _ = os.Create("/tmp/myga.ini")
	}
	file.Close()

	cfg, err := ini.Load("/tmp/myga.ini")
	if err != nil {
		panic(err)
	}

	if _, err := cfg.GetSection(item); err != nil {
		fmt.Println(item + " is not Exist!!!!")
		os.Exit(1)
	}

	cfg.DeleteSection(item)
	cfg.SaveTo("/tmp/myga.ini")
	fmt.Println("ok!!!")
	return nil
}

func modify(item string, secret string) error {
	file,er:=os.Open("/tmp/myga.ini")

	if er!=nil && os.IsNotExist(er){
		file, _ = os.Create("/tmp/myga.ini")
	}
	file.Close()

	cfg, err := ini.Load("/tmp/myga.ini")
	if err != nil {
		panic(err)
	}

	if _, err := cfg.GetSection(item); err != nil {
		fmt.Println(item + " is not Exist!!!!")
		os.Exit(1)
	}

	cfg.Section(item).Key("secret").SetValue(secret)
	cfg.SaveTo("/tmp/myga.ini")
	return nil
}


func main() {
	Cli()
}



