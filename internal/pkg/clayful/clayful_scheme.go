package clayful

type ClayfulProduct struct {
    ID          string    `json:"_id"`
    Name        string    `json:"name"`
    Price       PriceInfo `json:"price"`
    // Discount    Discount  `json:"discount"`
    // Shipping    Shipping  `json:"shipping"`
    // Rating      Rating    `json:"rating"`
    // Bundled     bool      `json:"bundled"`
    // Available   bool      `json:"available"`
    // Brand       Brand     `json:"brand"`
    // Thumbnail   Image     `json:"thumbnail"`
    // // TaxCategories omitted as it's an empty array in the example
    // TotalReview Count     `json:"totalReview"`
    // Collections []Collection `json:"collections"`
    // Catalogs    []Catalog `json:"catalogs"`
    // Options     []Option  `json:"options"`
    Variants    []Variant `json:"variants"`

}

type PriceInfo struct {
    Original Price `json:"original"`
    Sale     Price `json:"sale"`
}

type Price struct {
    Raw          int    `json:"raw"`
    ConvertedRaw int    `json:"convertedRaw"`
    Formatted    string `json:"formatted"`
    Converted    string `json:"converted"`
}

type Variation struct {
    Value    string `json:"value"`
    Priority int    `json:"priority"`
    ID       string `json:"_id"`
}


type Variant struct {
    Price     PriceInfo    `json:"price"`
    Types     []VariantType `json:"types"`
    ID        string       `json:"_id"`
}

type VariantType struct {
    Option    OptionVariant `json:"option"`
    Variation Variation     `json:"variation"`
}

type OptionVariant struct {
    Name     string `json:"name"`
    Priority int    `json:"priority"`
    ID       string `json:"_id"`
}

