package types

type Package struct {
	PackageBasicData
}

type PackageBasicData struct {
	ID        string          `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	//Type      pkg.Type        `json:"type"`
	//FoundBy   string          `json:"foundBy"`
	//Locations []file.Location `json:"locations"`
	//Licenses  licenses        `json:"licenses"`
	//Language  pkg.Language    `json:"language"`
	//CPEs      cpes            `json:"cpes"`
	PURL string `json:"purl"`
}

type PacakgeResponse struct {
}