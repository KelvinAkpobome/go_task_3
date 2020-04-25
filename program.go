package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type SocialMedia interface {
	Feed() []string
}

func TextFile(u SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		return errors.New("error: " + err.Error())
	}

	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("error: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}

func JSONFile(u SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()

	if err != nil {
		return errors.New("error: " + err.Error())
	}

	i := 1
	md := u.Feed()
	cd := make(map[int][]string)
	for _, value := range md {
		cd[i] = append(cd[i], value)
		i++
	}

	b, err := json.MarshalIndent(cd, "\n", "")

	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}

	bytesWritten, err := f.Write(b)

	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}

	fmt.Printf("wrote %d bytes\n", bytesWritten)

	return nil
}

func XMLFile(u SocialMedia, filename string) error {
	x, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer x.Close()

	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	md := u.Feed()
	b, err := xml.MarshalIndent(md, "\n", "\n")
	if err != nil {
		return errors.New("error: " + err.Error())
	}

	bytesWritten, err := x.Write(b)

	if err != nil {
		return errors.New("an error: " + err.Error())
	}
	fmt.Printf("wrote %d bytes\n", bytesWritten)

	return nil
}

//
func YAMLFile(u SocialMedia, filename string) error {
	q, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer q.Close()

	if err != nil {
		return errors.New("error: " + err.Error())
	}

	md := u.Feed()
	b, err := yaml.Marshal(md)
	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}

	bytesWritten, err := q.Write(b)

	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}
	fmt.Printf("wrote %d bytes\n", bytesWritten)

	return nil
}
