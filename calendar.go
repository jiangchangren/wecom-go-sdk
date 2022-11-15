package wework

import (
	"encoding/json"
	"fmt"

	"github.com/go-laoji/wecom-go-sdk/internal"
)

type CreateCalendarReq struct {
	Calendar Calendar `json:"calendar"`
	AgentId  int      `json:"agentid"`
}

type Calendar struct {
	Organizer      string `json:"organizer"`
	SetAsDefault   int    `json:"set_as_default"`
	IsPublic       int    `json:"is_public"`
	IsCorpCalendar int    `json:"is_corp_calendar"`
	CalendarSimple
}

type CalendarSimple struct {
	CalId       string `json:"cal_id,omitempty"`
	Readonly    int    `json:"readonly"`
	Summary     string `json:"summary"`
	Color       string `json:"color"`
	Description string `json:"description"`
	Shares      []struct {
		Userid   string `json:"userid"`
		Readonly int    `json:"readonly"`
	} `json:"shares"`
	PublicRange struct {
		UserIds  []string `json:"userids"`
		PartyIds []int    `json:"partyids"`
	} `json:"public_range"`
}

type CreateCalendarResponse struct {
	internal.BizResponse
	CalId string `json:"cal_id"`
}

type UpdateCalendarReq struct {
	SkipPublicRange int            `json:"skip_public_range"`
	Calendar        CalendarSimple `json:"calendar"`
}

type GetCalendarReq struct {
	CalIdList []string `json:"cal_id_list"`
}

type GetCalendarResponse struct {
	internal.BizResponse
	CalendarList []Calendar `json:"calendar_list"`
}

type DeleteCalendarReq struct {
	CalId string `json:"cal_id"`
}

// CreateCalendar 创建日历
func (ww weWork) CreateCalendar(corpId uint, req *CreateCalendarReq) (resp CreateCalendarResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/oa/calendar/add?%s", queryParams.Encode()), req)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// UpdateCalendar 更新日历
func (ww weWork) UpdateCalendar(corpId uint, req *UpdateCalendarReq) (resp internal.BizResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/oa/calendar/update?%s", queryParams.Encode()), req)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// GetCalendars 获取日历
func (ww weWork) GetCalendars(corpId uint, req *GetCalendarReq) (resp GetCalendarResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/oa/calendar/get?%s", queryParams.Encode()), req)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// DeleteCalendar 删除日历
func (ww weWork) DeleteCalendar(corpId uint, req *DeleteCalendarReq) (resp internal.BizResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/oa/calendar/del?%s", queryParams.Encode()), req)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
