package champions

type Champion struct {
	ID      string   `bson:"_id,omitempty" json:"id,omitempty"`
	Key     string   `bson:"key,omitempty" json:"key,omitempty"`
	KeyInt  int      `bson:"key_int,omitempty" json:"key_int,omitempty"`
	URL     string   `bson:"url,omitempty" json:"url,omitempty"`
	Name    string   `bson:"name,omitempty" json:"name,omitempty"`
	Title   string   `bson:"title,omitempty" json:"title,omitempty"`
	Image   *Image   `bson:"image,omitempty" json:"image,omitempty"`
	Lore    string   `bson:"lore,omitempty" json:"lore,omitempty"`
	Blurb   string   `bson:"blurb,omitempty" json:"blurb,omitempty"`
	Tags    []string `bson:"tags,omitempty" json:"tags,omitempty"`
	Spells  []Spell  `bson:"spells,omitempty" json:"spells,omitempty"`
	Passive *Passive `bson:"passive,omitempty" json:"passive,omitempty"`
}

type Image struct {
	Full   string `bson:"full,omitempty" json:"full,omitempty"`
	Sprite string `bson:"sprite,omitempty" json:"sprite,omitempty"`
	Group  string `bson:"group,omitempty" json:"group,omitempty"`
	X      int    `bson:"x,omitempty" json:"x,omitempty"`
	Y      int    `bson:"y,omitempty" json:"y,omitempty"`
	W      int    `bson:"w,omitempty" json:"w,omitempty"`
	H      int    `bson:"h,omitempty" json:"h,omitempty"`
}

type Spell struct {
	ID           string              `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string              `bson:"name,omitempty" json:"name,omitempty"`
	Description  string              `bson:"description,omitempty" json:"description,omitempty"`
	Tooltip      string              `bson:"tooltip,omitempty" json:"tooltip,omitempty"`
	LevelTip     map[string][]string `bson:"leveltip,omitempty" json:"leveltip,omitempty"`
	MaxRank      float32             `bson:"maxrank,omitempty" json:"maxrank,omitempty"`
	Cooldown     []float32           `bson:"cooldown,omitempty" json:"cooldown,omitempty"`
	CooldownBurn string              `bson:"cooldownBurn,omitempty" json:"cooldownBurn,omitempty"`
	Cost         []int               `bson:"cost,omitempty" json:"cost,omitempty"`
	CostBurn     string              `bson:"costBurn,omitempty" json:"costBurn,omitempty"`
	Effect       [][]float32         `bson:"effect,omitempty" json:"effect,omitempty"`
	EffectBurn   []string            `bson:"effectBurn,omitempty" json:"effectBurn,omitempty"`
	Vars         []string            `bson:"vars,omitempty" json:"vars,omitempty"`
	CostType     string              `bson:"costType,omitempty" json:"costType,omitempty"`
	MaxAmmo      string              `bson:"maxammo,omitempty" json:"maxammo,omitempty"`
	Range        []float32           `bson:"range,omitempty" json:"range,omitempty"`
	RangeBurn    string              `bson:"rangeBurn,omitempty" json:"rangeBurn,omitempty"`
	Image        *Image              `bson:"image,omitempty" json:"image,omitempty"`
	Resource     string              `bson:"resource,omitempty" json:"resource,omitempty"`
}

type Passive struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Image       *Image `bson:"image,omitempty" json:"image,omitempty"`
}
