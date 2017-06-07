// Coral Talk is a package to interact with talk API
//
// Coral System
//
// Talk is the service to get data into the coral system.  Each endpoint is documented in https://github.com/coralproject/talk/tree/master/docs
// In this package we are using:
//
package talk

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/coralproject/coral-talk-driver/pkg/webservice"
)

type Talk struct {
	endpoint string
}

// Connect initialization of coral talk connection.
func Connect(endpoint string) *Talk {

	talk := &Talk{
		endpoint: endpoint
	}

	return talk
}


/*
 * Using the Coral Talk GraphQL API
 */

// Get a comment by ID
func (c *Talk) GetComment(id string) ([]map[string]Interface{}, error) {

	data := {
		id
	}

	if _, ok := c.endpoint ; ok {

		d, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Error %v", err)
		}

		userAgent := fmt.Sprintf("Sponge Publisher")
		result, err = webservice.DoRequest(userAgent, methodPost, c.endpoint, bytes.NewBuffer(d))

		return result
	}
}