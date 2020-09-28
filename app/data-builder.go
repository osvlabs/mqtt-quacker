package app

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// DataBuilderConfig - Configuration of MQTT server
type DataBuilderConfig struct {
	Path string // Path - Data template file path
}

// DataBuilder - The databuilder class.
type DataBuilder struct {
	config   DataBuilderConfig
	template string
	slots    map[int]Slot
}

// NewDataBuilder - Create a new DataBuilder object
func NewDataBuilder(config DataBuilderConfig) DataBuilder {
	builder := DataBuilder{
		config: config,
	}
	err := builder.parse()
	if err != nil {
		panic(err)
	}
	return builder
}

// Close - Close the databuilder mission
func (b *DataBuilder) Close() {
}

// Parse - Parse the template JSON to initialize the builder
func (b *DataBuilder) parse() error {
	rawTemplate, err := ioutil.ReadFile(b.config.Path)
	if err != nil {
		return err
	}

	matcher, err := regexp.Compile("(\"q:.*\")")
	if err != nil {
		return err
	}

	slotCount := 0
	innerMatcher, err := regexp.Compile("\"q:(float):(.*),(.*)\"")
	if err != nil {
		return err
	}
	slots := make(map[int]Slot)

	parsedTemplate := matcher.ReplaceAllFunc(rawTemplate, func(bytes []byte) []byte {
		slotCount = slotCount + 1
		slots[slotCount] = Slot{
			count:    slotCount,
			seed:     int(rand.Float32() * 100),
			provider: b.parseProvider(slotCount, innerMatcher, bytes),
		}

		return []byte("${" + strconv.Itoa(slotCount) + "}")
	})

	// fmt.Printf("parsed %s\n", string(parsedTemplate))

	b.template = string(parsedTemplate)
	b.slots = slots

	return nil
}

// parseProvider - Parse the slot to get value provider function.
func (b *DataBuilder) parseProvider(slotCount int, innerMatcher *regexp.Regexp, bytes []byte) Provider {
	result := innerMatcher.FindAllSubmatch(bytes, 10)
	valueType := string(result[0][1])
	minValue, err := strconv.ParseFloat(string(result[0][2]), 64)
	if err != nil {
		panic(err)
	}
	maxValue, err := strconv.ParseFloat(string(result[0][3]), 64)
	if err != nil {
		panic(err)
	}

	return func() string {
		if valueType == "float" {
			return strconv.FormatFloat(rand.Float64()*(maxValue-minValue)+minValue, 'f', 10, 64)
		}
		return "unknown"
	}
}

// Make - Make a payload
func (b *DataBuilder) Make() (string, error) {
	matcher, err := regexp.Compile(`\${\d*}`)
	if err != nil {
		return "", err
	}

	payload := matcher.ReplaceAllStringFunc(b.template, func(slotCountString string) string {
		slotCount, err := strconv.Atoi(strings.Trim(slotCountString, "${}"))
		if err != nil {
			panic(err)
		}
		return b.slots[slotCount].provider()
	})

	return payload, nil
}

type Slot struct {
	count    int
	seed     int
	provider Provider
}

type Provider func() string
