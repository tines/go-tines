package tines

import (
	"context"

	"fmt"
	"time"
)

// AgentService handles fields for the Tines instance / API.
type AgentService struct {
	client *Client
}

// Agent structure
type Agent struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	Options struct {
		Mode     string `json:"mode"`
		Lookback string `json:"lookback"`
		Path     string `json:"path"`
	} `json:"options"`
	Name               string      `json:"name"`
	Schedule           interface{} `json:"schedule"`
	EventsCount        int         `json:"events_count"`
	LastCheckAt        interface{} `json:"last_check_at"`
	LastReceiveAt      time.Time   `json:"last_receive_at"`
	LastCheckedEventID int         `json:"last_checked_event_id"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	LastWebRequestAt   interface{} `json:"last_web_request_at"`
	KeepEventsFor      int         `json:"keep_events_for"`
	LastEventAt        time.Time   `json:"last_event_at"`
	LastErrorLogAt     interface{} `json:"last_error_log_at"`
	Disabled           bool        `json:"disabled"`
	GUID               string      `json:"guid"`
	StoryID            int         `json:"story_id"`
}

// Options structure
type Options struct {
	Secret string `json:"secret"`
	Verbs  string `json:"verbs"`
}

// AgentReq structure
type AgentReq struct {
	Type          string   `json:"type"`
	Name          string   `json:"name"`
	StoryID       int      `json:"story_id"`
	KeepEventsFor int      `json:"keep_events_for"`
	SourceIds     []string `json:"source_ids"`
	ReceiverIds   []string `json:"receiver_ids"`
	Options       Options  `json:"options"`
}

// GetWithContext returns an agent for the given agent key.
func (s *AgentService) GetWithContext(ctx context.Context, agentID string) (*Agent, *Response, error) {
	apiEndpoint := fmt.Sprintf("agents/%s", agentID)
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	agent := new(Agent)
	resp, err := s.client.Do(req, agent)
	if err != nil {
		return nil, resp, err
	}

	return agent, resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *AgentService) Get(agentID string) (*Agent, *Response, error) {
	return s.GetWithContext(context.Background(), agentID)
}

// DeleteWithContext deletes an agent.
func (s *AgentService) DeleteWithContext(ctx context.Context, agentID string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("agents/%s", agentID)
	req, err := s.client.NewRequestWithContext(ctx, "DELETE", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Delete wraps GetWithContext using the background context.
func (s *AgentService) Delete(agentID string) (*Response, error) {
	return s.DeleteWithContext(context.Background(), agentID)
}

// CreateWithContext creates an agent.
func (s *AgentService) CreateWithContext(ctx context.Context, agentreq *AgentReq) (*Agent, *Response, error) {
	apiEndpoint := fmt.Sprintf("agents")
	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, agentreq)
	if err != nil {
		return nil, nil, err
	}

	agent := new(Agent)
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}

	return agent, resp, err
}

// Create wraps CreateWithContext using the background context.
func (s *AgentService) Create(agentreq *AgentReq) (*Agent, *Response, error) {
	return s.CreateWithContext(context.Background(), agentreq)
}

// UpdateWithContext updates an agent for the given agent key.
func (s *AgentService) UpdateWithContext(ctx context.Context, agentID string, agentreq *AgentReq) (*Agent, *Response, error) {
	apiEndpoint := fmt.Sprintf("agents/%s", agentID)
	req, err := s.client.NewRequestWithContext(ctx, "PUT", apiEndpoint, agentreq)
	if err != nil {
		return nil, nil, err
	}

	agent := new(Agent)
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}

	return agent, resp, err
}

// Update wraps UpdateWithContext using the background context.
func (s *AgentService) Update(agentID string, agentreq *AgentReq) (*Agent, *Response, error) {
	return s.UpdateWithContext(context.Background(), agentID, agentreq)
}
