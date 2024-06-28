package common

import (
	"os"
	"path/filepath"

	"github.com/8VIM/keyboard_layout_calculator/common/internal/utils"
)

func InitCache() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(homeDir, ".config", "keyboard_layout_calculator", "cache")

	if err := utils.EnsureDirExists(dir); err != nil {
		return "", err
	}
	return dir, nil
}
