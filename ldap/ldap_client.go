package ldap_client

import (
	"log"

	"github.com/jtblin/go-ldap-client"
)

func Client_Authenticate(username string, password string) {
	client := &ldap.LDAPClient{
		Base:               "dc=example,dc=net",
		Host:               "ldaps.examplae.net",
		Port:               389,
		UseSSL:             true,
		BindDN:             "username",
		BindPassword:       "password",
		UserFilter:         "(sAMAccountName=%s)",
		InsecureSkipVerify: true,
		Attributes:         []string{"givenName", "sn", "mail", "sAMAccountName"},
	}
	// It is the responsibility of the caller to close the connection
	defer client.Close()

	ok, user, err := client.Authenticate(username, password)
	if err != nil {
		log.Fatalf("Error authenticating user %s: %+v", username, err)
	}
	if !ok {
		log.Fatalf("Authenticating failed for user %s", username)
	}
	log.Printf("User: %+v", user)
	//  return user
}
