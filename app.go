package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx        context.Context
	gamePath   string
	configPath string
}

// ModType represents the type of mod loader required
type ModType string

const (
	ModTypeMelonLoader ModType = "MelonLoader"
	ModTypeBepInEx     ModType = "BepInEx"
	ModTypeBoth        ModType = "Both"
)

// ModCategory represents different categories of mods
type ModCategory string

const (
	ModCategoryPerformance   ModCategory = "Performance"
	ModCategoryQualityOfLife ModCategory = "Quality of Life"
	ModCategoryContent       ModCategory = "Content"
	ModCategoryOverhaul      ModCategory = "Overhaul"
	ModCategoryOther         ModCategory = "Other"
)

// Default paths to check for Desktop Mate installation
var defaultPaths = []string{
	`C:\SteamLibrary\steamapps\common\Desktop Mate`,
	`C:\Program Files (x86)\Steam\steamapps\common\Desktop Mate`,
}

// Config represents the application configuration
type Config struct {
	GamePath string `json:"gamePath"`
}

// Mod represents basic mod information
type Mod struct {
	Name string `json:"name"`
	Repo string `json:"repo"`
}

// Sponsor represents sponsorship information
type Sponsor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Locale represents localization information
type Locale struct {
	Default   string   `json:"default"`
	Supported []string `json:"supported"`
}

// ModMeta represents detailed metadata for a mod
type ModMeta struct {
	Description string        `json:"description"`
	Author      string        `json:"author"`
	Category    []ModCategory `json:"category"`
	Type        ModType       `json:"type"`
	Version     string        `json:"version"`
	Requires    []string      `json:"requires,omitempty"`
	Sponsor     *Sponsor      `json:"sponsor,omitempty"`
	Locale      *Locale       `json:"locale,omitempty"`
}

// APIResponse represents the structure of the API response
type APIResponse struct {
	Status int       `json:"status"`
	Data   []ModData `json:"data"`
}

// ModData extends Mod with additional API-specific fields
type ModData struct {
	Name      string   `json:"name"`
	Repo      string   `json:"repo"`
	Downloads int      `json:"downloads"`
	Views     int      `json:"views"`
	Meta      ModMeta  `json:"meta"`
	Locale    Locale   `json:"locale"`
	Featured  bool     `json:"featured"`
	ID        string   `json:"_id"`
	Version   int      `json:"version"`
	Requires  []string `json:"requires"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	if err := a.LoadConfig(); err != nil {
		fmt.Printf("Warning: Failed to load config: %v\n", err)
	}
	if a.gamePath == "" {
		if err := a.DetectGamePath(); err != nil {
			fmt.Printf("Warning: %v\n", err)
		}
	}
}

// getConfigPath returns the path to the config file
func (a *App) getConfigPath() (string, error) {
	if a.configPath != "" {
		return a.configPath, nil
	}

	appData, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get config directory: %w", err)
	}

	configDir := filepath.Join(appData, "DesktopMateModManager")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	a.configPath = filepath.Join(configDir, "config.json")
	return a.configPath, nil
}

// SaveConfig saves the current configuration to disk
func (a *App) SaveConfig() error {
	configPath, err := a.getConfigPath()
	if err != nil {
		return err
	}

	config := Config{
		GamePath: a.gamePath,
	}

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// LoadConfig loads the configuration from disk
func (a *App) LoadConfig() error {
	configPath, err := a.getConfigPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	if config.GamePath != "" && isValidGamePath(config.GamePath) {
		a.gamePath = config.GamePath
	}

	return nil
}

// DetectGamePath attempts to find Desktop Mate installation
func (a *App) DetectGamePath() error {
	if a.gamePath != "" {
		if isValidGamePath(a.gamePath) {
			return nil
		}
		return fmt.Errorf("invalid custom game path: %s", a.gamePath)
	}

	for _, path := range defaultPaths {
		if isValidGamePath(path) {
			a.gamePath = path
			if err := a.SaveConfig(); err != nil {
				return fmt.Errorf("failed to save detected game path: %w", err)
			}
			return nil
		}
	}

	return fmt.Errorf("Desktop Mate installation not found in default locations")
}

// isValidGamePath checks if the given path contains Desktop Mate
func isValidGamePath(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !info.IsDir() {
		return false
	}

	requiredFiles := []string{
		"DesktopMate.exe",
		"UnityPlayer.dll",
		"DesktopMate_Data",
	}

	for _, file := range requiredFiles {
		filePath := filepath.Join(path, file)
		if _, err := os.Stat(filePath); err != nil {
			fmt.Printf("Missing required file: %s\n", file) // Debug log
			return false
		}
	}

	return true
}

// SetGamePath opens a dialog to select a directory for game path
func (a *App) SetGamePath(ctx context.Context) error {
	path, err := runtime.OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{
		Title: "Select Game Directory",
	})
	if err != nil {
		return err
	}

	if path == "" {
		return fmt.Errorf("no path selected")
	}

	return a.SetCustomGamePath(path)
}

// SetCustomGamePath sets a custom path for the game installation
func (a *App) SetCustomGamePath(path string) error {
	cleanPath := filepath.Clean(path)

	if !isValidGamePath(cleanPath) {
		return fmt.Errorf("invalid game path: %s", cleanPath)
	}

	a.gamePath = cleanPath

	if err := a.SaveConfig(); err != nil {
		return fmt.Errorf("failed to save game path: %w", err)
	}

	return nil
}

// GetGamePath returns the current game path
func (a *App) GetGamePath() string {
	return a.gamePath
}

// ResetGamePath clears the saved game path
func (a *App) ResetGamePath() error {
	a.gamePath = ""
	return a.SaveConfig()
}

// FetchMods retrieves mods from the API
func (a *App) FetchMods() ([]ModData, error) {
	resp, err := http.Get("https://api.dskt.cc/mods")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mods: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	return apiResponse.Data, nil
}

// GetAvailableMods returns all available mods from the API
func (a *App) GetAvailableMods() []ModData {
	mods, err := a.FetchMods()
	if err != nil {
		fmt.Printf("Error fetching mods: %v\n", err)
		return nil
	}
	return mods
}
