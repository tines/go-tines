package tines

import (
	"context"

	"fmt"
)

// AnnotationService handles fields for the Tines instance / API.
type AnnotationService struct {
	client *Client
}

// Annotation structure
type Annotation struct {
	ID       int                    `json:"id" structs:"id,omitempty"`
	StoryID  int                    `json:"story_id" structs:"story_id"`
	Content  string                 `json:"content" structs:"content,omitempty"`
	Position map[string]interface{} `json:"position" structs:"position,omitempty"`
}

// GetWithContext returns an annotation for the given annotation key.
func (s *AnnotationService) GetWithContext(ctx context.Context, annotationID int) (*Annotation, *Response, error) {
	apiEndpoint := fmt.Sprintf("annotations/%v", annotationID)
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	annotation := new(Annotation)
	resp, err := s.client.Do(req, annotation)
	if err != nil {
		return nil, resp, err
	}

	return annotation, resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *AnnotationService) Get(annotationID int) (*Annotation, *Response, error) {
	return s.GetWithContext(context.Background(), annotationID)
}

// DeleteWithContext deletes an annotation.
func (s *AnnotationService) DeleteWithContext(ctx context.Context, annotationID int) (*Response, error) {
	apiEndpoint := fmt.Sprintf("annotations/%v", annotationID)
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
func (s *AnnotationService) Delete(annotationID int) (*Response, error) {
	return s.DeleteWithContext(context.Background(), annotationID)
}

// CreateWithContext creates an annotation.
func (s *AnnotationService) CreateWithContext(ctx context.Context, annotation *Annotation) (*Annotation, *Response, error) {
	apiEndpoint := fmt.Sprintf("annotations")
	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, annotation)
	if err != nil {
		return nil, nil, err
	}

	annotationresp := new(Annotation)
	resp, err := s.client.Do(req, annotationresp)
	if err != nil {
		return nil, resp, err
	}

	return annotationresp, resp, err
}

// Create wraps CreateWithContext using the background context.
func (s *AnnotationService) Create(annotation *Annotation) (*Annotation, *Response, error) {
	return s.CreateWithContext(context.Background(), annotation)
}

// UpdateWithContext updates an annotation for the given annotation key.
func (s *AnnotationService) UpdateWithContext(ctx context.Context, annotationID int, annotation *Annotation) (*Annotation, *Response, error) {
	apiEndpoint := fmt.Sprintf("annotations/%v", annotationID)
	req, err := s.client.NewRequestWithContext(ctx, "PUT", apiEndpoint, annotation)
	if err != nil {
		return nil, nil, err
	}

	annotationresp := new(Annotation)
	resp, err := s.client.Do(req, annotationresp)
	if err != nil {
		return nil, resp, err
	}

	return annotationresp, resp, err
}

// Update wraps UpdateWithContext using the background context.
func (s *AnnotationService) Update(annotationID int, annotation *Annotation) (*Annotation, *Response, error) {
	return s.UpdateWithContext(context.Background(), annotationID, annotation)
}
