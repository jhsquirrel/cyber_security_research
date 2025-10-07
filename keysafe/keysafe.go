package main

import (
	"flag"
	"fmt"
	"os"
	"errors"
	"log"
)

func help() {
    fmt.Println("The -user AND -site MUST be specified")
    fmt.Println(os.Args[0] + " -user=<username> -site=<sitename>")
    fmt.Println("-pass OR -get MUST be specified")
    fmt.Println("e.g. " + os.Args[0] + " -user me@me.com -site mysite.com -get")
    log.Fatal("invalid arguments")
}

func createstore_and_pass_journey(KEYPASSFILE string, username *string, sitename *string) (error) {
    fmt.Println("Enter keysafe password")
    _, err, keysafepass := getHiddenUserInput()
    if err != nil {
	return err
    }
    fmt.Println("Enter password for -user and -site to be saved in keystore")
    _, err, userpass := getHiddenUserInput()
    if err != nil {
	return err
    }
    // we now have the keysafepass, the userpass and can create a keystore
    createKeyStore(KEYPASSFILE)
    var storeData []keyuser
    storeData = append(storeData, keyuser{User: *username, Pass: userpass, Site: *sitename})
    setDataStore(KEYPASSFILE, storeData, keysafepass)
    return nil
}

func getpass_journey(KEYPASSFILE string, username *string, sitename *string) (string, error) {
    fmt.Println("Attempting to get key for " + *username + " @ " + *sitename)
    // get the password for the keystore
    fmt.Println("Enter keysafe password")
    _, err, keysafepass := getHiddenUserInput()
    check(err)
    data := getDataStore(KEYPASSFILE, keysafepass)
    for _, element := range data {
        if element.User == *username && element.Site == *sitename {
	    return element.Pass, nil
        }
    }
    return "", errors.New("No matches found")
}

func main() {
    KEYPASSFILE := "./keypass.store"

    var newpass = flag.Bool("pass", false, "add/update a new password")
    var sitename = flag.String("site", "", "associate a password with a site")
    var username = flag.String("user", "", "associate a password and site with a user")
    var getpass = flag.Bool("get", false, "get the stored password")

    flag.Parse()

    if *sitename == "" && *username == "" {
        help()
    } else if *newpass == false && *getpass == false {
        help()
    } else if *newpass == true && *getpass == true {
        help()
    }

    // check if a keysafe exists, if not, potentially create it
    _, err := os.Stat(KEYPASSFILE)
    if err == nil {
        // keyfile exists
	if *getpass == true {
	    var rv_pass, err = getpass_journey(KEYPASSFILE, username, sitename)
	    if err != nil {
		log.Fatal(err)
	    }
	    fmt.Println("Password : ", rv_pass)
	} else if *newpass == true {
	    // get the password for the keystore
	    fmt.Println("enter password for keystore")
	    _, err, keysafepass := getHiddenUserInput()
            check(err)
	    data := getDataStore(KEYPASSFILE, keysafepass)
	    var found = false
            for index, element := range data {
		if element.User == *username && element.Site == *sitename {
		    found = true
                    fmt.Println("The Username and Sitename already exist with a password!")
                    fmt.Println("Do you with to update the password?")
                    var _, _, YN = getYNUserInput()
		    if YN == "N" {
			log.Fatal("ok. exiting")
		    } else {
                        // update password
			fmt.Println("Update pass here")
                        _, err, userpass := getHiddenUserInput()
                        if err != nil {
			    log.Fatal(err)
                        }
			data[index].Pass = userpass
                        setDataStore(KEYPASSFILE, data, keysafepass)
		    }
                }
            }
	    if found == false {
		// there is no entry for this user and site already
                fmt.Println("Checking if a password for " + *username + " @ " + *sitename + " exists")
	        fmt.Println("Enter password for this user and site")
	        _, err, userpass := getHiddenUserInput()
                if err != nil {
		    log.Fatal(err)
                }
                data = append(data, keyuser{User: *username, Pass: userpass, Site: *sitename})
                setDataStore(KEYPASSFILE, data, keysafepass)
            }
	}
    } else {
        // keyfile does not exist
	if *getpass == true {
            fmt.Println("The keyfile does not exist - you must add passwords to it first")
	} else if *newpass == true {
	    err := createstore_and_pass_journey(KEYPASSFILE, username, sitename)
	    check(err)
	}
    }
}
