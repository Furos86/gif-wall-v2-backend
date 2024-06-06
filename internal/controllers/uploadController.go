package controllers

import (
	"gw-backend/internal/services/interactionLayer"
)

type UploadController struct {
	connetionManager interactionLayer.ConnectionManager
}
