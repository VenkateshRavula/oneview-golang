package ov

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type SNMPv3Trap struct {
	Category           string        `json:"category,omitempty"`
	Created            string        `json:"created,omitempty"`
	DestinationAddress string        `json:"destinationAddress,omitempty"`
	ETAG               string        `json:"eTag,omitempty"`
	EngineID           string        `json:"engineId,omitempty"`
	ID                 string        `json:"id,omitempty"`
	Modified           string        `json:"modified,omitempty"`
	Port               int           `json:"port,omitempty"`
	TrapType           string        `json:"traptype,omitempty"`
	Type               string        `json:"type,omitempty"`
	URI                utils.Nstring `json:"uri,omitempty"`
	UserID             string        `json:"userId,omitempty"`
	UserURI            string        `json:"userUri,omitempty"`
}

type SNMPv3TrapList struct {
	Count       int           `json:"count,omitempty"`
	Category    string        `json:"category,omitempty"`
	Created     string        `json:"created,omitempty"`
	ETAG        string        `json:"eTag,omitempty"`
	Members     []SNMPv3Trap  `json:"members,omitempty"`     // "members":[]
	Modified    string        `json:"modified,omitempty"`    // "modified":null"
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Type        string        `json:"type,omitempty"`
	URI         utils.Nstring `json:"uri,omitempty"`
}

type ValidateSNMPv3Address struct {
	DestinationAddress   string          `json:"destinationAddress,omitempty"`
	ExistingDestinations []utils.Nstring `json:"existingDestinations,omitempty"`
}

func (c *OVClient) CreateSNMPv3TrapDestinations(trapOption SNMPv3Trap) error {
	log.Infof("Initializing creation of SNMPv3 Trap Destinations for %s.", trapOption.UserID)
	var (
		uri      = "/rest/appliance/snmpv3-trap-forwarding/destinations"
		trapdata SNMPv3Trap
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, trapOption)
	data, err := c.RestAPICall(rest.POST, uri, trapOption)
	if err != nil {
		log.Errorf("Error submitting new snmpv3 trap destinations request: %s", err)
		return err
	}

	log.Debugf("Response New SNMPv3 Trap Destinations %s", data)
	if err := json.Unmarshal(data, &trapdata); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) GetSNMPv3TrapDestinations(filter string, sort string, start string, count string) (SNMPv3TrapList, error) {
	var (
		uri      = "/rest/appliance/snmpv3-trap-forwarding/destinations"
		q        = make(map[string]interface{})
		traplist SNMPv3TrapList
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	if start != "" {
		q["start"] = start
	}

	if count != "" {
		q["count"] = count
	}
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return traplist, err
	}

	log.Debugf("Gettraplist %s", data)
	if err := json.Unmarshal(data, &traplist); err != nil {
		return traplist, err
	}
	return traplist, nil

}

func (c *OVClient) GetSNMPv3TrapDestinationsById(id string) (SNMPv3Trap, error) {
	var (
		uri    = "/rest/appliance/snmpv3-trap-forwarding/destinations/" + id
		trapId SNMPv3Trap
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return trapId, err
	}

	log.Debugf("GetLocalelist %s", data)
	if err := json.Unmarshal(data, &trapId); err != nil {
		return trapId, err
	}
	return trapId, nil
}

func (c *OVClient) ValidateDestinationAddress(validateId ValidateSNMPv3Address) error {
	var (
		uri = "/rest/appliance/snmpv3-trap-forwarding/destinations/validation"
		t   *Task
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, validateId)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, validateId)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting validating destination address request: %s", err)
		return err
	}

	log.Debugf("Validated Destination Address")
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}
	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) UpdateSNMPv3TrapDestinations(updateOption SNMPv3Trap) (SNMPv3Trap, error) {
	if updateOption.ID == "" {
		fmt.Println("The ID field is not set")
		return updateOption, errors.New("ID field is empty")
	}

	var (
		uri            = "/rest/appliance/snmpv3-trap-forwarding/destinations/" + updateOption.ID
		updateResponse SNMPv3Trap
		t              *Task
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, updateOption)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, updateOption)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update SNMPv3 Trap Destination request: %s", err)
		return updateResponse, err
	}

	log.Debugf("Response SNMPv3 Trap Destination %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return updateResponse, err
	}

	if err := json.Unmarshal(data, &updateResponse); err != nil {
		return updateResponse, err
	}

	return updateResponse, nil
}

func (c *OVClient) DeleteSNMPv3TrapDestinations(id string) error {
	var (
		trap SNMPv3Trap
		err  error
		t    *Task
		uri  string
	)

	trap, err = c.GetSNMPv3TrapDestinationsById(id)
	if err != nil {
		return err
	}
	if trap.ID != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", trap.URI, trap)
		log.Debugf("task -> %+v", t)
		uri = trap.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete snmp trap destination request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete snmp trap destination %s", data)
		if err := json.Unmarshal([]byte(data), &t); err != nil {
			t.TaskIsDone = true
			log.Errorf("Error with task un-marshal: %s", err)
			return err
		}
		err = t.Wait()
		if err != nil {
			return err
		}
		return nil
	} else {
		log.Infof("SNMPv3TrapDestination could not be found to delete, %s, skipping delete ...", id)
	}
	return nil
}
