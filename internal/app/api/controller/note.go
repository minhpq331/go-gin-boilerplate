package controller

import (
	"net/http"

	"go-gin-boilerplate/internal/common"
	"go-gin-boilerplate/internal/model"
	"go-gin-boilerplate/internal/service"

	"github.com/gin-gonic/gin"
)

// NoteController handles HTTP requests for notes
type NoteController struct {
	noteService *service.NoteService
}

// NewNoteController creates a new note controller
func NewNoteController(noteService *service.NoteService) *NoteController {
	return &NoteController{
		noteService: noteService,
	}
}

// CreateNote handles note creation
// @Summary Create a new note
// @Description Create a new note with the provided details
// @Tags notes
// @Accept json
// @Produce json
// @Param note body model.Note true "Note object"
// @Success 201 {object} model.Note
// @Failure 400 {object} map[string]string
// @Router /notes [post]
func (c *NoteController) CreateNote(ctx *gin.Context) {
	var note model.Note
	if err := ctx.ShouldBindJSON(&note); err != nil {
		panic(common.NewAPIError(err.Error(), common.RequestValidationError, http.StatusBadRequest, err.Error()))
	}

	if err := c.noteService.CreateNote(ctx, &note); err != nil {
		panic(common.NewAPIError("Failed to create the note", common.RequestValidationError, http.StatusBadRequest, err.Error()))
	}

	ctx.JSON(http.StatusCreated, common.SendResponse(note))
}

// GetNote handles retrieving a note by ID
// @Summary Get a note by ID
// @Description Get a note by its ID
// @Tags notes
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} model.Note
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [get]
func (c *NoteController) GetNote(ctx *gin.Context) {
	id := ctx.Param("id")
	note, err := c.noteService.GetNote(ctx, id)
	if err != nil {
		panic(common.NewAPIError("Failed to get the note", common.RequestValidationError, http.StatusBadRequest, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SendResponse(note))
}

// GetAllNotes handles retrieving all notes
// @Summary Get all notes
// @Description Get a list of all notes
// @Tags notes
// @Accept json
// @Produce json
// @Success 200 {array} model.Note
// @Router /notes [get]
func (c *NoteController) GetAllNotes(ctx *gin.Context) {
	notes, err := c.noteService.GetAllNotes(ctx)
	if err != nil {
		panic(common.NewAPIError("Failed to get notes", common.RequestValidationError, http.StatusBadRequest, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SendResponse(notes))
}

// UpdateNote handles note updates
// @Summary Update a note
// @Description Update a note with the provided details
// @Tags notes
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Param note body model.Note true "Note object"
// @Success 200 {object} model.Note
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [put]
func (c *NoteController) UpdateNote(ctx *gin.Context) {
	id := ctx.Param("id")
	var note model.Note
	if err := ctx.ShouldBindJSON(&note); err != nil {
		panic(common.NewAPIError(err.Error(), common.RequestValidationError, http.StatusBadRequest, err.Error()))

	}

	existingNote, err := c.noteService.GetNote(ctx, id)
	if err != nil {
		panic(common.NewAPIError("Failed to get the note", common.RequestValidationError, http.StatusBadRequest, err.Error()))
	}

	note.ID = existingNote.ID
	if err := c.noteService.UpdateNote(ctx, &note); err != nil {
		panic(common.NewAPIError("Failed to update the note", common.RequestValidationError, http.StatusBadRequest, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SendResponse(note))
}

// DeleteNote handles note deletion
// @Summary Delete a note
// @Description Delete a note by its ID
// @Tags notes
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [delete]
func (c *NoteController) DeleteNote(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.noteService.DeleteNote(ctx, id); err != nil {
		panic(common.NewAPIError("Failed to delete the note", common.RequestValidationError, http.StatusBadRequest, err.Error()))
	}

	ctx.Status(http.StatusNoContent)
}
