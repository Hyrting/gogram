package react

import (
	"github.com/hyrting/gogram"
)

type Reaction = gogram.Reaction

func React(r ...Reaction) gogram.Reactions {
	return gogram.Reactions{Reactions: r}
}

// Currently available emojis.
var (
	ThumbUp                   = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👍"}
	ThumbDown                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👎"}
	Heart                     = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "❤"}
	Fire                      = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🔥"}
	HeartEyes                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😍"}
	ClappingHands             = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👏"}
	GrinningFace              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😁"}
	ThinkingFace              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤔"}
	ExplodingHead             = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤯"}
	ScreamingFace             = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😱"}
	SwearingFace              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤬"}
	CryingFace                = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😢"}
	PartyPopper               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🎉"}
	StarStruck                = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤩"}
	VomitingFace              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤮"}
	PileOfPoo                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "💩"}
	PrayingHands              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🙏"}
	OkHand                    = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👌"}
	DoveOfPeace               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🕊"}
	ClownFace                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤡"}
	YawningFace               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🥱"}
	WoozyFace                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🥴"}
	Whale                     = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🐳"}
	HeartOnFire               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "❤‍🔥"}
	MoonFace                  = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🌚"}
	HotDog                    = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🌭"}
	HundredPoints             = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "💯"}
	RollingOnTheFloorLaughing = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤣"}
	Lightning                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "⚡"}
	Banana                    = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🍌"}
	Trophy                    = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🏆"}
	BrokenHeart               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "💔"}
	FaceWithRaisedEyebrow     = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤨"}
	NeutralFace               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😐"}
	Strawberry                = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🍓"}
	Champagne                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🍾"}
	KissMark                  = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "💋"}
	MiddleFinger              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🖕"}
	EvilFace                  = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😈"}
	SleepingFace              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😴"}
	LoudlyCryingFace          = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😭"}
	NerdFace                  = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤓"}
	Ghost                     = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👻"}
	Engineer                  = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👨‍💻"}
	Eyes                      = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👀"}
	JackOLantern              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🎃"}
	NoMonkey                  = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🙈"}
	SmilingFaceWithHalo       = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😇"}
	FearfulFace               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😨"}
	Handshake                 = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤝"}
	WritingHand               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "✍"}
	HuggingFace               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤗"}
	Brain                     = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🫡"}
	SantaClaus                = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🎅"}
	ChristmasTree             = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🎄"}
	Snowman                   = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "☃"}
	NailPolish                = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "💅"}
	ZanyFace                  = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤪"}
	Moai                      = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🗿"}
	Cool                      = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🆒"}
	HeartWithArrow            = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "💘"}
	HearMonkey                = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🙉"}
	Unicorn                   = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🦄"}
	FaceBlowingKiss           = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😘"}
	Pill                      = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "💊"}
	SpeaklessMonkey           = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🙊"}
	Sunglasses                = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😎"}
	AlienMonster              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "👾"}
	ManShrugging              = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤷‍♂️"}
	PersonShrugging           = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤷"}
	WomanShrugging            = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "🤷‍♀️"}
	PoutingFace               = Reaction{Type: gogram.ReactionTypeEmoji, Emoji: "😡"}
)
