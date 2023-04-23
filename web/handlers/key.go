package handlers

import (
	"fmt"
	"github.com/hutaochu/hello-hutao/internal/types/res"
	"github.com/hutaochu/hello-hutao/pkg/utils"
	"github.com/hutaochu/hello-hutao/pkg/utils/logger"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GetPublicKey
// @BasePath /
// @Summary get public key
// @Schemes
// @Description get public key from server, use this to encryption
// @Tags hello-hutao
// @Accept json
// @Produce json
// @Success 200 {object} res.GetPublicKeyResponse
// @Router /public_key [get]
func GetPublicKey(c *gin.Context) {
	p, err := filepath.Abs("cert/public.pem")
	if err != nil {
		logger.Error(c, err, fmt.Sprintf("get execute path err"))
		utils.SetError(c, err)
		return
	}
	f, err := os.ReadFile(p)
	if err != nil {
		logger.Error(c, err, fmt.Sprintf("read public key err"))
		utils.SetError(c, err)
		return
	}
	r := res.GetPublicKeyResponse{
		Key: utils.ToString(f),
	}
	utils.SetResponse(c, r)
}
