package services

import (
	"context"

	"github.com/dstm45/vacances/pkg/database"
	"github.com/google/uuid"
)

type NoteService struct {
	db *database.Queries
}

func NewNoteService(db database.Queries) *NoteService {
	pool, err := database.Connection()
	if err != nil {
		return nil
	}
	noteService := NoteService{}
	return &noteService
}

func (s NoteService) CreateNote(ctx context.Context, params database.CreateNoteParams) (*database.Note, error) {
	note, err := s.db.CreateNote(ctx, params)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (s NoteService) DeleteNote(ctx context.Context, id uuid.UUID) error {
	err := s.db.DeleteNote(ctx, id)
	return err
}

func (s NoteService) GetNote(ctx context.Context, id uuid.UUID) (*database.Note, error) {
	note, err := s.db.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	return &note, nil
}
