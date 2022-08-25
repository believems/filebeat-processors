package impala_profile

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/beats/v7/libbeat/common"
	"time"
)

type ImpalaProfile struct {
	Timestamp time.Time `json:"timestamp"`
	QueryId   string    `json:"queryId"`
	Profile   string    `json:"profile"`
}

func (profile *ImpalaProfile) StringMap() (common.MapStr, error) {
	var stringMap common.MapStr
	data, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &stringMap)
	if err != nil {
		return nil, err
	}
	return stringMap, nil
}

func (profile *ImpalaProfile) MarshalJSON() ([]byte, error) {
	type Alias ImpalaProfile
	return json.Marshal(&struct {
		Timestamp string `json:"timestamp"`
		*Alias
	}{
		Timestamp: profile.Timestamp.Format(time.RFC3339),
		Alias:     (*Alias)(profile),
	})
}

func (profile *ImpalaProfile) String() string {
	timeStr := profile.Timestamp.Format(time.RFC3339)
	return fmt.Sprintf("%s %s %s", timeStr, profile.QueryId, profile.Profile)
}