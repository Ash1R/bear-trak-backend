// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"time"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/BusRoute
type BusRoute struct {
	ID         int                 `json:"id"`
	SortIdx    int                 `json:"sortIdx"`
	Name       string              `json:"name"`
	Code       string              `json:"code"`
	Color      string              `json:"color"`
	Directions []BusRouteDirection `json:"directions"`
	Vehicles   []Vehicle           `json:"vehicles"`
}

// GetID returns the value of ID.
func (s *BusRoute) GetID() int {
	return s.ID
}

// GetSortIdx returns the value of SortIdx.
func (s *BusRoute) GetSortIdx() int {
	return s.SortIdx
}

// GetName returns the value of Name.
func (s *BusRoute) GetName() string {
	return s.Name
}

// GetCode returns the value of Code.
func (s *BusRoute) GetCode() string {
	return s.Code
}

// GetColor returns the value of Color.
func (s *BusRoute) GetColor() string {
	return s.Color
}

// GetDirections returns the value of Directions.
func (s *BusRoute) GetDirections() []BusRouteDirection {
	return s.Directions
}

// GetVehicles returns the value of Vehicles.
func (s *BusRoute) GetVehicles() []Vehicle {
	return s.Vehicles
}

// SetID sets the value of ID.
func (s *BusRoute) SetID(val int) {
	s.ID = val
}

// SetSortIdx sets the value of SortIdx.
func (s *BusRoute) SetSortIdx(val int) {
	s.SortIdx = val
}

// SetName sets the value of Name.
func (s *BusRoute) SetName(val string) {
	s.Name = val
}

// SetCode sets the value of Code.
func (s *BusRoute) SetCode(val string) {
	s.Code = val
}

// SetColor sets the value of Color.
func (s *BusRoute) SetColor(val string) {
	s.Color = val
}

// SetDirections sets the value of Directions.
func (s *BusRoute) SetDirections(val []BusRouteDirection) {
	s.Directions = val
}

// SetVehicles sets the value of Vehicles.
func (s *BusRoute) SetVehicles(val []Vehicle) {
	s.Vehicles = val
}

// Ref: #/components/schemas/BusRouteDirection
type BusRouteDirection struct {
	ID        BusRouteDirectionID          `json:"id"`
	Polylines []string                     `json:"polylines"`
	Stops     []BusRouteDirectionStopsItem `json:"stops"`
}

// GetID returns the value of ID.
func (s *BusRouteDirection) GetID() BusRouteDirectionID {
	return s.ID
}

// GetPolylines returns the value of Polylines.
func (s *BusRouteDirection) GetPolylines() []string {
	return s.Polylines
}

// GetStops returns the value of Stops.
func (s *BusRouteDirection) GetStops() []BusRouteDirectionStopsItem {
	return s.Stops
}

// SetID sets the value of ID.
func (s *BusRouteDirection) SetID(val BusRouteDirectionID) {
	s.ID = val
}

// SetPolylines sets the value of Polylines.
func (s *BusRouteDirection) SetPolylines(val []string) {
	s.Polylines = val
}

// SetStops sets the value of Stops.
func (s *BusRouteDirection) SetStops(val []BusRouteDirectionStopsItem) {
	s.Stops = val
}

// Ref: #/components/schemas/BusRouteDirectionID
type BusRouteDirectionID string

const (
	BusRouteDirectionIDInbound     BusRouteDirectionID = "inbound"
	BusRouteDirectionIDOutbound    BusRouteDirectionID = "outbound"
	BusRouteDirectionIDUnspecified BusRouteDirectionID = "unspecified"
)

// AllValues returns all BusRouteDirectionID values.
func (BusRouteDirectionID) AllValues() []BusRouteDirectionID {
	return []BusRouteDirectionID{
		BusRouteDirectionIDInbound,
		BusRouteDirectionIDOutbound,
		BusRouteDirectionIDUnspecified,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s BusRouteDirectionID) MarshalText() ([]byte, error) {
	switch s {
	case BusRouteDirectionIDInbound:
		return []byte(s), nil
	case BusRouteDirectionIDOutbound:
		return []byte(s), nil
	case BusRouteDirectionIDUnspecified:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *BusRouteDirectionID) UnmarshalText(data []byte) error {
	switch BusRouteDirectionID(data) {
	case BusRouteDirectionIDInbound:
		*s = BusRouteDirectionIDInbound
		return nil
	case BusRouteDirectionIDOutbound:
		*s = BusRouteDirectionIDOutbound
		return nil
	case BusRouteDirectionIDUnspecified:
		*s = BusRouteDirectionIDUnspecified
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type BusRouteDirectionStopsItem struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// GetID returns the value of ID.
func (s *BusRouteDirectionStopsItem) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *BusRouteDirectionStopsItem) GetName() string {
	return s.Name
}

// GetLongitude returns the value of Longitude.
func (s *BusRouteDirectionStopsItem) GetLongitude() float64 {
	return s.Longitude
}

// GetLatitude returns the value of Latitude.
func (s *BusRouteDirectionStopsItem) GetLatitude() float64 {
	return s.Latitude
}

// SetID sets the value of ID.
func (s *BusRouteDirectionStopsItem) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *BusRouteDirectionStopsItem) SetName(val string) {
	s.Name = val
}

// SetLongitude sets the value of Longitude.
func (s *BusRouteDirectionStopsItem) SetLongitude(val float64) {
	s.Longitude = val
}

// SetLatitude sets the value of Latitude.
func (s *BusRouteDirectionStopsItem) SetLatitude(val float64) {
	s.Latitude = val
}

// Ref: #/components/schemas/Eatery
type Eatery struct {
	ID             int                    `json:"id"`
	Name           string                 `json:"name"`
	NameShort      string                 `json:"nameShort"`
	ImagePath      string                 `json:"imagePath"`
	Latitude       float64                `json:"latitude"`
	Longitude      float64                `json:"longitude"`
	Location       string                 `json:"location"`
	Hours          []Hours                `json:"hours"`
	Region         EateryRegion           `json:"region"`
	PayMethods     []EateryPayMethodsItem `json:"payMethods"`
	Categories     []EateryCategoriesItem `json:"categories"`
	NextWeekEvents EateryNextWeekEvents   `json:"nextWeekEvents"`
	AllWeekMenu    []EateryMenuCategory   `json:"allWeekMenu"`
}

// GetID returns the value of ID.
func (s *Eatery) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *Eatery) GetName() string {
	return s.Name
}

// GetNameShort returns the value of NameShort.
func (s *Eatery) GetNameShort() string {
	return s.NameShort
}

// GetImagePath returns the value of ImagePath.
func (s *Eatery) GetImagePath() string {
	return s.ImagePath
}

// GetLatitude returns the value of Latitude.
func (s *Eatery) GetLatitude() float64 {
	return s.Latitude
}

// GetLongitude returns the value of Longitude.
func (s *Eatery) GetLongitude() float64 {
	return s.Longitude
}

// GetLocation returns the value of Location.
func (s *Eatery) GetLocation() string {
	return s.Location
}

// GetHours returns the value of Hours.
func (s *Eatery) GetHours() []Hours {
	return s.Hours
}

// GetRegion returns the value of Region.
func (s *Eatery) GetRegion() EateryRegion {
	return s.Region
}

// GetPayMethods returns the value of PayMethods.
func (s *Eatery) GetPayMethods() []EateryPayMethodsItem {
	return s.PayMethods
}

// GetCategories returns the value of Categories.
func (s *Eatery) GetCategories() []EateryCategoriesItem {
	return s.Categories
}

// GetNextWeekEvents returns the value of NextWeekEvents.
func (s *Eatery) GetNextWeekEvents() EateryNextWeekEvents {
	return s.NextWeekEvents
}

// GetAllWeekMenu returns the value of AllWeekMenu.
func (s *Eatery) GetAllWeekMenu() []EateryMenuCategory {
	return s.AllWeekMenu
}

// SetID sets the value of ID.
func (s *Eatery) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Eatery) SetName(val string) {
	s.Name = val
}

// SetNameShort sets the value of NameShort.
func (s *Eatery) SetNameShort(val string) {
	s.NameShort = val
}

// SetImagePath sets the value of ImagePath.
func (s *Eatery) SetImagePath(val string) {
	s.ImagePath = val
}

// SetLatitude sets the value of Latitude.
func (s *Eatery) SetLatitude(val float64) {
	s.Latitude = val
}

// SetLongitude sets the value of Longitude.
func (s *Eatery) SetLongitude(val float64) {
	s.Longitude = val
}

// SetLocation sets the value of Location.
func (s *Eatery) SetLocation(val string) {
	s.Location = val
}

// SetHours sets the value of Hours.
func (s *Eatery) SetHours(val []Hours) {
	s.Hours = val
}

// SetRegion sets the value of Region.
func (s *Eatery) SetRegion(val EateryRegion) {
	s.Region = val
}

// SetPayMethods sets the value of PayMethods.
func (s *Eatery) SetPayMethods(val []EateryPayMethodsItem) {
	s.PayMethods = val
}

// SetCategories sets the value of Categories.
func (s *Eatery) SetCategories(val []EateryCategoriesItem) {
	s.Categories = val
}

// SetNextWeekEvents sets the value of NextWeekEvents.
func (s *Eatery) SetNextWeekEvents(val EateryNextWeekEvents) {
	s.NextWeekEvents = val
}

// SetAllWeekMenu sets the value of AllWeekMenu.
func (s *Eatery) SetAllWeekMenu(val []EateryMenuCategory) {
	s.AllWeekMenu = val
}

type EateryCategoriesItem string

const (
	EateryCategoriesItemConvenienceStore EateryCategoriesItem = "convenienceStore"
	EateryCategoriesItemCafe             EateryCategoriesItem = "cafe"
	EateryCategoriesItemDiningRoom       EateryCategoriesItem = "diningRoom"
	EateryCategoriesItemCoffeeShop       EateryCategoriesItem = "coffeeShop"
	EateryCategoriesItemCart             EateryCategoriesItem = "cart"
	EateryCategoriesItemFoodCourt        EateryCategoriesItem = "foodCourt"
)

// AllValues returns all EateryCategoriesItem values.
func (EateryCategoriesItem) AllValues() []EateryCategoriesItem {
	return []EateryCategoriesItem{
		EateryCategoriesItemConvenienceStore,
		EateryCategoriesItemCafe,
		EateryCategoriesItemDiningRoom,
		EateryCategoriesItemCoffeeShop,
		EateryCategoriesItemCart,
		EateryCategoriesItemFoodCourt,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s EateryCategoriesItem) MarshalText() ([]byte, error) {
	switch s {
	case EateryCategoriesItemConvenienceStore:
		return []byte(s), nil
	case EateryCategoriesItemCafe:
		return []byte(s), nil
	case EateryCategoriesItemDiningRoom:
		return []byte(s), nil
	case EateryCategoriesItemCoffeeShop:
		return []byte(s), nil
	case EateryCategoriesItemCart:
		return []byte(s), nil
	case EateryCategoriesItemFoodCourt:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *EateryCategoriesItem) UnmarshalText(data []byte) error {
	switch EateryCategoriesItem(data) {
	case EateryCategoriesItemConvenienceStore:
		*s = EateryCategoriesItemConvenienceStore
		return nil
	case EateryCategoriesItemCafe:
		*s = EateryCategoriesItemCafe
		return nil
	case EateryCategoriesItemDiningRoom:
		*s = EateryCategoriesItemDiningRoom
		return nil
	case EateryCategoriesItemCoffeeShop:
		*s = EateryCategoriesItemCoffeeShop
		return nil
	case EateryCategoriesItemCart:
		*s = EateryCategoriesItemCart
		return nil
	case EateryCategoriesItemFoodCourt:
		*s = EateryCategoriesItemFoodCourt
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/EateryEvent
type EateryEvent struct {
	Start          time.Time            `json:"start"`
	End            time.Time            `json:"end"`
	MenuCategories []EateryMenuCategory `json:"menuCategories"`
}

// GetStart returns the value of Start.
func (s *EateryEvent) GetStart() time.Time {
	return s.Start
}

// GetEnd returns the value of End.
func (s *EateryEvent) GetEnd() time.Time {
	return s.End
}

// GetMenuCategories returns the value of MenuCategories.
func (s *EateryEvent) GetMenuCategories() []EateryMenuCategory {
	return s.MenuCategories
}

// SetStart sets the value of Start.
func (s *EateryEvent) SetStart(val time.Time) {
	s.Start = val
}

// SetEnd sets the value of End.
func (s *EateryEvent) SetEnd(val time.Time) {
	s.End = val
}

// SetMenuCategories sets the value of MenuCategories.
func (s *EateryEvent) SetMenuCategories(val []EateryMenuCategory) {
	s.MenuCategories = val
}

// Ref: #/components/schemas/EateryMenuCategory
type EateryMenuCategory struct {
	Name  string                        `json:"name"`
	Items []EateryMenuCategoryItemsItem `json:"items"`
}

// GetName returns the value of Name.
func (s *EateryMenuCategory) GetName() string {
	return s.Name
}

// GetItems returns the value of Items.
func (s *EateryMenuCategory) GetItems() []EateryMenuCategoryItemsItem {
	return s.Items
}

// SetName sets the value of Name.
func (s *EateryMenuCategory) SetName(val string) {
	s.Name = val
}

// SetItems sets the value of Items.
func (s *EateryMenuCategory) SetItems(val []EateryMenuCategoryItemsItem) {
	s.Items = val
}

type EateryMenuCategoryItemsItem struct {
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
}

// GetName returns the value of Name.
func (s *EateryMenuCategoryItemsItem) GetName() string {
	return s.Name
}

// GetHealthy returns the value of Healthy.
func (s *EateryMenuCategoryItemsItem) GetHealthy() bool {
	return s.Healthy
}

// SetName sets the value of Name.
func (s *EateryMenuCategoryItemsItem) SetName(val string) {
	s.Name = val
}

// SetHealthy sets the value of Healthy.
func (s *EateryMenuCategoryItemsItem) SetHealthy(val bool) {
	s.Healthy = val
}

type EateryNextWeekEvents struct {
	Monday    []EateryEvent `json:"monday"`
	Tuesday   []EateryEvent `json:"tuesday"`
	Wednesday []EateryEvent `json:"wednesday"`
	Thursday  []EateryEvent `json:"thursday"`
	Friday    []EateryEvent `json:"friday"`
	Saturday  []EateryEvent `json:"saturday"`
	Sunday    []EateryEvent `json:"sunday"`
}

// GetMonday returns the value of Monday.
func (s *EateryNextWeekEvents) GetMonday() []EateryEvent {
	return s.Monday
}

// GetTuesday returns the value of Tuesday.
func (s *EateryNextWeekEvents) GetTuesday() []EateryEvent {
	return s.Tuesday
}

// GetWednesday returns the value of Wednesday.
func (s *EateryNextWeekEvents) GetWednesday() []EateryEvent {
	return s.Wednesday
}

// GetThursday returns the value of Thursday.
func (s *EateryNextWeekEvents) GetThursday() []EateryEvent {
	return s.Thursday
}

// GetFriday returns the value of Friday.
func (s *EateryNextWeekEvents) GetFriday() []EateryEvent {
	return s.Friday
}

// GetSaturday returns the value of Saturday.
func (s *EateryNextWeekEvents) GetSaturday() []EateryEvent {
	return s.Saturday
}

// GetSunday returns the value of Sunday.
func (s *EateryNextWeekEvents) GetSunday() []EateryEvent {
	return s.Sunday
}

// SetMonday sets the value of Monday.
func (s *EateryNextWeekEvents) SetMonday(val []EateryEvent) {
	s.Monday = val
}

// SetTuesday sets the value of Tuesday.
func (s *EateryNextWeekEvents) SetTuesday(val []EateryEvent) {
	s.Tuesday = val
}

// SetWednesday sets the value of Wednesday.
func (s *EateryNextWeekEvents) SetWednesday(val []EateryEvent) {
	s.Wednesday = val
}

// SetThursday sets the value of Thursday.
func (s *EateryNextWeekEvents) SetThursday(val []EateryEvent) {
	s.Thursday = val
}

// SetFriday sets the value of Friday.
func (s *EateryNextWeekEvents) SetFriday(val []EateryEvent) {
	s.Friday = val
}

// SetSaturday sets the value of Saturday.
func (s *EateryNextWeekEvents) SetSaturday(val []EateryEvent) {
	s.Saturday = val
}

// SetSunday sets the value of Sunday.
func (s *EateryNextWeekEvents) SetSunday(val []EateryEvent) {
	s.Sunday = val
}

type EateryPayMethodsItem string

const (
	EateryPayMethodsItemSwipes        EateryPayMethodsItem = "swipes"
	EateryPayMethodsItemBigRedBucks   EateryPayMethodsItem = "bigRedBucks"
	EateryPayMethodsItemCash          EateryPayMethodsItem = "cash"
	EateryPayMethodsItemDigitalWallet EateryPayMethodsItem = "digitalWallet"
	EateryPayMethodsItemCreditCard    EateryPayMethodsItem = "creditCard"
	EateryPayMethodsItemCornellCard   EateryPayMethodsItem = "cornellCard"
)

// AllValues returns all EateryPayMethodsItem values.
func (EateryPayMethodsItem) AllValues() []EateryPayMethodsItem {
	return []EateryPayMethodsItem{
		EateryPayMethodsItemSwipes,
		EateryPayMethodsItemBigRedBucks,
		EateryPayMethodsItemCash,
		EateryPayMethodsItemDigitalWallet,
		EateryPayMethodsItemCreditCard,
		EateryPayMethodsItemCornellCard,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s EateryPayMethodsItem) MarshalText() ([]byte, error) {
	switch s {
	case EateryPayMethodsItemSwipes:
		return []byte(s), nil
	case EateryPayMethodsItemBigRedBucks:
		return []byte(s), nil
	case EateryPayMethodsItemCash:
		return []byte(s), nil
	case EateryPayMethodsItemDigitalWallet:
		return []byte(s), nil
	case EateryPayMethodsItemCreditCard:
		return []byte(s), nil
	case EateryPayMethodsItemCornellCard:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *EateryPayMethodsItem) UnmarshalText(data []byte) error {
	switch EateryPayMethodsItem(data) {
	case EateryPayMethodsItemSwipes:
		*s = EateryPayMethodsItemSwipes
		return nil
	case EateryPayMethodsItemBigRedBucks:
		*s = EateryPayMethodsItemBigRedBucks
		return nil
	case EateryPayMethodsItemCash:
		*s = EateryPayMethodsItemCash
		return nil
	case EateryPayMethodsItemDigitalWallet:
		*s = EateryPayMethodsItemDigitalWallet
		return nil
	case EateryPayMethodsItemCreditCard:
		*s = EateryPayMethodsItemCreditCard
		return nil
	case EateryPayMethodsItemCornellCard:
		*s = EateryPayMethodsItemCornellCard
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type EateryRegion string

const (
	EateryRegionCentral EateryRegion = "central"
	EateryRegionWest    EateryRegion = "west"
	EateryRegionNorth   EateryRegion = "north"
	EateryRegionUnknown EateryRegion = "unknown"
)

// AllValues returns all EateryRegion values.
func (EateryRegion) AllValues() []EateryRegion {
	return []EateryRegion{
		EateryRegionCentral,
		EateryRegionWest,
		EateryRegionNorth,
		EateryRegionUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s EateryRegion) MarshalText() ([]byte, error) {
	switch s {
	case EateryRegionCentral:
		return []byte(s), nil
	case EateryRegionWest:
		return []byte(s), nil
	case EateryRegionNorth:
		return []byte(s), nil
	case EateryRegionUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *EateryRegion) UnmarshalText(data []byte) error {
	switch EateryRegion(data) {
	case EateryRegionCentral:
		*s = EateryRegionCentral
		return nil
	case EateryRegionWest:
		*s = EateryRegionWest
		return nil
	case EateryRegionNorth:
		*s = EateryRegionNorth
		return nil
	case EateryRegionUnknown:
		*s = EateryRegionUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/Gym
type Gym struct {
	ID                  int                          `json:"id"`
	Name                string                       `json:"name"`
	ImagePath           string                       `json:"imagePath"`
	Latitude            float64                      `json:"latitude"`
	Longitude           float64                      `json:"longitude"`
	Hours               []Hours                      `json:"hours"`
	Facilities          []GymFacilitiesItem          `json:"facilities"`
	EquipmentCategories []GymEquipmentCategoriesItem `json:"equipmentCategories"`
	Capacity            NilGymCapacity               `json:"capacity"`
}

// GetID returns the value of ID.
func (s *Gym) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *Gym) GetName() string {
	return s.Name
}

// GetImagePath returns the value of ImagePath.
func (s *Gym) GetImagePath() string {
	return s.ImagePath
}

// GetLatitude returns the value of Latitude.
func (s *Gym) GetLatitude() float64 {
	return s.Latitude
}

// GetLongitude returns the value of Longitude.
func (s *Gym) GetLongitude() float64 {
	return s.Longitude
}

// GetHours returns the value of Hours.
func (s *Gym) GetHours() []Hours {
	return s.Hours
}

// GetFacilities returns the value of Facilities.
func (s *Gym) GetFacilities() []GymFacilitiesItem {
	return s.Facilities
}

// GetEquipmentCategories returns the value of EquipmentCategories.
func (s *Gym) GetEquipmentCategories() []GymEquipmentCategoriesItem {
	return s.EquipmentCategories
}

// GetCapacity returns the value of Capacity.
func (s *Gym) GetCapacity() NilGymCapacity {
	return s.Capacity
}

// SetID sets the value of ID.
func (s *Gym) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Gym) SetName(val string) {
	s.Name = val
}

// SetImagePath sets the value of ImagePath.
func (s *Gym) SetImagePath(val string) {
	s.ImagePath = val
}

// SetLatitude sets the value of Latitude.
func (s *Gym) SetLatitude(val float64) {
	s.Latitude = val
}

// SetLongitude sets the value of Longitude.
func (s *Gym) SetLongitude(val float64) {
	s.Longitude = val
}

// SetHours sets the value of Hours.
func (s *Gym) SetHours(val []Hours) {
	s.Hours = val
}

// SetFacilities sets the value of Facilities.
func (s *Gym) SetFacilities(val []GymFacilitiesItem) {
	s.Facilities = val
}

// SetEquipmentCategories sets the value of EquipmentCategories.
func (s *Gym) SetEquipmentCategories(val []GymEquipmentCategoriesItem) {
	s.EquipmentCategories = val
}

// SetCapacity sets the value of Capacity.
func (s *Gym) SetCapacity(val NilGymCapacity) {
	s.Capacity = val
}

type GymCapacity struct {
	Count       int       `json:"count"`
	Percentage  NilInt    `json:"percentage"`
	LastUpdated time.Time `json:"lastUpdated"`
}

// GetCount returns the value of Count.
func (s *GymCapacity) GetCount() int {
	return s.Count
}

// GetPercentage returns the value of Percentage.
func (s *GymCapacity) GetPercentage() NilInt {
	return s.Percentage
}

// GetLastUpdated returns the value of LastUpdated.
func (s *GymCapacity) GetLastUpdated() time.Time {
	return s.LastUpdated
}

// SetCount sets the value of Count.
func (s *GymCapacity) SetCount(val int) {
	s.Count = val
}

// SetPercentage sets the value of Percentage.
func (s *GymCapacity) SetPercentage(val NilInt) {
	s.Percentage = val
}

// SetLastUpdated sets the value of LastUpdated.
func (s *GymCapacity) SetLastUpdated(val time.Time) {
	s.LastUpdated = val
}

type GymEquipmentCategoriesItem struct {
	CategoryType GymEquipmentCategoriesItemCategoryType `json:"categoryType"`
	Items        []string                               `json:"items"`
}

// GetCategoryType returns the value of CategoryType.
func (s *GymEquipmentCategoriesItem) GetCategoryType() GymEquipmentCategoriesItemCategoryType {
	return s.CategoryType
}

// GetItems returns the value of Items.
func (s *GymEquipmentCategoriesItem) GetItems() []string {
	return s.Items
}

// SetCategoryType sets the value of CategoryType.
func (s *GymEquipmentCategoriesItem) SetCategoryType(val GymEquipmentCategoriesItemCategoryType) {
	s.CategoryType = val
}

// SetItems sets the value of Items.
func (s *GymEquipmentCategoriesItem) SetItems(val []string) {
	s.Items = val
}

type GymEquipmentCategoriesItemCategoryType string

const (
	GymEquipmentCategoriesItemCategoryTypeTreadmills  GymEquipmentCategoriesItemCategoryType = "treadmills"
	GymEquipmentCategoriesItemCategoryTypeEllipticals GymEquipmentCategoriesItemCategoryType = "ellipticals"
	GymEquipmentCategoriesItemCategoryTypeRowing      GymEquipmentCategoriesItemCategoryType = "rowing"
	GymEquipmentCategoriesItemCategoryTypeBike        GymEquipmentCategoriesItemCategoryType = "bike"
	GymEquipmentCategoriesItemCategoryTypeLifting     GymEquipmentCategoriesItemCategoryType = "lifting"
	GymEquipmentCategoriesItemCategoryTypeMachines    GymEquipmentCategoriesItemCategoryType = "machines"
	GymEquipmentCategoriesItemCategoryTypeFreeWeights GymEquipmentCategoriesItemCategoryType = "freeWeights"
	GymEquipmentCategoriesItemCategoryTypeMisc        GymEquipmentCategoriesItemCategoryType = "misc"
)

// AllValues returns all GymEquipmentCategoriesItemCategoryType values.
func (GymEquipmentCategoriesItemCategoryType) AllValues() []GymEquipmentCategoriesItemCategoryType {
	return []GymEquipmentCategoriesItemCategoryType{
		GymEquipmentCategoriesItemCategoryTypeTreadmills,
		GymEquipmentCategoriesItemCategoryTypeEllipticals,
		GymEquipmentCategoriesItemCategoryTypeRowing,
		GymEquipmentCategoriesItemCategoryTypeBike,
		GymEquipmentCategoriesItemCategoryTypeLifting,
		GymEquipmentCategoriesItemCategoryTypeMachines,
		GymEquipmentCategoriesItemCategoryTypeFreeWeights,
		GymEquipmentCategoriesItemCategoryTypeMisc,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s GymEquipmentCategoriesItemCategoryType) MarshalText() ([]byte, error) {
	switch s {
	case GymEquipmentCategoriesItemCategoryTypeTreadmills:
		return []byte(s), nil
	case GymEquipmentCategoriesItemCategoryTypeEllipticals:
		return []byte(s), nil
	case GymEquipmentCategoriesItemCategoryTypeRowing:
		return []byte(s), nil
	case GymEquipmentCategoriesItemCategoryTypeBike:
		return []byte(s), nil
	case GymEquipmentCategoriesItemCategoryTypeLifting:
		return []byte(s), nil
	case GymEquipmentCategoriesItemCategoryTypeMachines:
		return []byte(s), nil
	case GymEquipmentCategoriesItemCategoryTypeFreeWeights:
		return []byte(s), nil
	case GymEquipmentCategoriesItemCategoryTypeMisc:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *GymEquipmentCategoriesItemCategoryType) UnmarshalText(data []byte) error {
	switch GymEquipmentCategoriesItemCategoryType(data) {
	case GymEquipmentCategoriesItemCategoryTypeTreadmills:
		*s = GymEquipmentCategoriesItemCategoryTypeTreadmills
		return nil
	case GymEquipmentCategoriesItemCategoryTypeEllipticals:
		*s = GymEquipmentCategoriesItemCategoryTypeEllipticals
		return nil
	case GymEquipmentCategoriesItemCategoryTypeRowing:
		*s = GymEquipmentCategoriesItemCategoryTypeRowing
		return nil
	case GymEquipmentCategoriesItemCategoryTypeBike:
		*s = GymEquipmentCategoriesItemCategoryTypeBike
		return nil
	case GymEquipmentCategoriesItemCategoryTypeLifting:
		*s = GymEquipmentCategoriesItemCategoryTypeLifting
		return nil
	case GymEquipmentCategoriesItemCategoryTypeMachines:
		*s = GymEquipmentCategoriesItemCategoryTypeMachines
		return nil
	case GymEquipmentCategoriesItemCategoryTypeFreeWeights:
		*s = GymEquipmentCategoriesItemCategoryTypeFreeWeights
		return nil
	case GymEquipmentCategoriesItemCategoryTypeMisc:
		*s = GymEquipmentCategoriesItemCategoryTypeMisc
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type GymFacilitiesItem struct {
	FacilityType GymFacilitiesItemFacilityType `json:"facilityType"`
	Name         string                        `json:"name"`
}

// GetFacilityType returns the value of FacilityType.
func (s *GymFacilitiesItem) GetFacilityType() GymFacilitiesItemFacilityType {
	return s.FacilityType
}

// GetName returns the value of Name.
func (s *GymFacilitiesItem) GetName() string {
	return s.Name
}

// SetFacilityType sets the value of FacilityType.
func (s *GymFacilitiesItem) SetFacilityType(val GymFacilitiesItemFacilityType) {
	s.FacilityType = val
}

// SetName sets the value of Name.
func (s *GymFacilitiesItem) SetName(val string) {
	s.Name = val
}

type GymFacilitiesItemFacilityType string

const (
	GymFacilitiesItemFacilityTypePool       GymFacilitiesItemFacilityType = "pool"
	GymFacilitiesItemFacilityTypeBasketball GymFacilitiesItemFacilityType = "basketball"
	GymFacilitiesItemFacilityTypeBowling    GymFacilitiesItemFacilityType = "bowling"
	GymFacilitiesItemFacilityTypeUnknown    GymFacilitiesItemFacilityType = "unknown"
)

// AllValues returns all GymFacilitiesItemFacilityType values.
func (GymFacilitiesItemFacilityType) AllValues() []GymFacilitiesItemFacilityType {
	return []GymFacilitiesItemFacilityType{
		GymFacilitiesItemFacilityTypePool,
		GymFacilitiesItemFacilityTypeBasketball,
		GymFacilitiesItemFacilityTypeBowling,
		GymFacilitiesItemFacilityTypeUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s GymFacilitiesItemFacilityType) MarshalText() ([]byte, error) {
	switch s {
	case GymFacilitiesItemFacilityTypePool:
		return []byte(s), nil
	case GymFacilitiesItemFacilityTypeBasketball:
		return []byte(s), nil
	case GymFacilitiesItemFacilityTypeBowling:
		return []byte(s), nil
	case GymFacilitiesItemFacilityTypeUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *GymFacilitiesItemFacilityType) UnmarshalText(data []byte) error {
	switch GymFacilitiesItemFacilityType(data) {
	case GymFacilitiesItemFacilityTypePool:
		*s = GymFacilitiesItemFacilityTypePool
		return nil
	case GymFacilitiesItemFacilityTypeBasketball:
		*s = GymFacilitiesItemFacilityTypeBasketball
		return nil
	case GymFacilitiesItemFacilityTypeBowling:
		*s = GymFacilitiesItemFacilityTypeBowling
		return nil
	case GymFacilitiesItemFacilityTypeUnknown:
		*s = GymFacilitiesItemFacilityTypeUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Hours
type Hours struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// GetStart returns the value of Start.
func (s *Hours) GetStart() time.Time {
	return s.Start
}

// GetEnd returns the value of End.
func (s *Hours) GetEnd() time.Time {
	return s.End
}

// SetStart sets the value of Start.
func (s *Hours) SetStart(val time.Time) {
	s.Start = val
}

// SetEnd sets the value of End.
func (s *Hours) SetEnd(val time.Time) {
	s.End = val
}

// NewNilGymCapacity returns new NilGymCapacity with value set to v.
func NewNilGymCapacity(v GymCapacity) NilGymCapacity {
	return NilGymCapacity{
		Value: v,
	}
}

// NilGymCapacity is nullable GymCapacity.
type NilGymCapacity struct {
	Value GymCapacity
	Null  bool
}

// SetTo sets value to v.
func (o *NilGymCapacity) SetTo(v GymCapacity) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilGymCapacity) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *NilGymCapacity) SetToNull() {
	o.Null = true
	var v GymCapacity
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilGymCapacity) Get() (v GymCapacity, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilGymCapacity) Or(d GymCapacity) GymCapacity {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewNilInt returns new NilInt with value set to v.
func NewNilInt(v int) NilInt {
	return NilInt{
		Value: v,
	}
}

// NilInt is nullable int.
type NilInt struct {
	Value int
	Null  bool
}

// SetTo sets value to v.
func (o *NilInt) SetTo(v int) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilInt) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *NilInt) SetToNull() {
	o.Null = true
	var v int
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilInt) Get() (v int, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *NilString) SetToNull() {
	o.Null = true
	var v string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Vehicle
type Vehicle struct {
	ID            int                 `json:"id"`
	RouteId       int                 `json:"routeId"`
	Name          string              `json:"name"`
	DirectionId   BusRouteDirectionID `json:"directionId"`
	Heading       int                 `json:"heading"`
	Latitude      float64             `json:"latitude"`
	Longitude     float64             `json:"longitude"`
	DisplayStatus string              `json:"displayStatus"`
	NextStop      string              `json:"nextStop"`
	LastStop      NilString           `json:"lastStop"`
	LastUpdated   time.Time           `json:"lastUpdated"`
}

// GetID returns the value of ID.
func (s *Vehicle) GetID() int {
	return s.ID
}

// GetRouteId returns the value of RouteId.
func (s *Vehicle) GetRouteId() int {
	return s.RouteId
}

// GetName returns the value of Name.
func (s *Vehicle) GetName() string {
	return s.Name
}

// GetDirectionId returns the value of DirectionId.
func (s *Vehicle) GetDirectionId() BusRouteDirectionID {
	return s.DirectionId
}

// GetHeading returns the value of Heading.
func (s *Vehicle) GetHeading() int {
	return s.Heading
}

// GetLatitude returns the value of Latitude.
func (s *Vehicle) GetLatitude() float64 {
	return s.Latitude
}

// GetLongitude returns the value of Longitude.
func (s *Vehicle) GetLongitude() float64 {
	return s.Longitude
}

// GetDisplayStatus returns the value of DisplayStatus.
func (s *Vehicle) GetDisplayStatus() string {
	return s.DisplayStatus
}

// GetNextStop returns the value of NextStop.
func (s *Vehicle) GetNextStop() string {
	return s.NextStop
}

// GetLastStop returns the value of LastStop.
func (s *Vehicle) GetLastStop() NilString {
	return s.LastStop
}

// GetLastUpdated returns the value of LastUpdated.
func (s *Vehicle) GetLastUpdated() time.Time {
	return s.LastUpdated
}

// SetID sets the value of ID.
func (s *Vehicle) SetID(val int) {
	s.ID = val
}

// SetRouteId sets the value of RouteId.
func (s *Vehicle) SetRouteId(val int) {
	s.RouteId = val
}

// SetName sets the value of Name.
func (s *Vehicle) SetName(val string) {
	s.Name = val
}

// SetDirectionId sets the value of DirectionId.
func (s *Vehicle) SetDirectionId(val BusRouteDirectionID) {
	s.DirectionId = val
}

// SetHeading sets the value of Heading.
func (s *Vehicle) SetHeading(val int) {
	s.Heading = val
}

// SetLatitude sets the value of Latitude.
func (s *Vehicle) SetLatitude(val float64) {
	s.Latitude = val
}

// SetLongitude sets the value of Longitude.
func (s *Vehicle) SetLongitude(val float64) {
	s.Longitude = val
}

// SetDisplayStatus sets the value of DisplayStatus.
func (s *Vehicle) SetDisplayStatus(val string) {
	s.DisplayStatus = val
}

// SetNextStop sets the value of NextStop.
func (s *Vehicle) SetNextStop(val string) {
	s.NextStop = val
}

// SetLastStop sets the value of LastStop.
func (s *Vehicle) SetLastStop(val NilString) {
	s.LastStop = val
}

// SetLastUpdated sets the value of LastUpdated.
func (s *Vehicle) SetLastUpdated(val time.Time) {
	s.LastUpdated = val
}
