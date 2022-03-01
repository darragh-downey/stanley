package app

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/darragh-downey/stanley/pkg/model"
)

func TestConcurrentParser(t *testing.T) {
	tests := []struct {
		name    string
		stream  []byte
		want    []model.StanleyResponse
		wantErr bool
	}{
		{
			name:    "empty case",
			stream:  []byte(""),
			want:    []model.StanleyResponse{},
			wantErr: false,
		},
		{
			name: "single case",
			stream: []byte(
				`{
						"payload":[
							{
								"country": "AUS",
								"description": "Join the most dynamic TV judging panel Australia has ever seen as they uncover the next breed of superstars every Sunday night. UK comedy royalty Dawn French, international pop superstar Geri Halliwell, (in) famous Aussie straight-talking radio jock Kyle Sandilands, and chart -topping former AGT alumni Timomatic.",
								"drm": true,
								"episodeCount": 0,
								"genre": "Reality",
								"image": {
									"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/AGT.jpg"
								},
								"language": "English",
								"nextEpisode": {
									"channel": null,
									"channelLogo": "http://catchup.ninemsn.com.au/img/player/Ch9_new_logo.gif",
									"date": null,
									"html": "Next episode airs:<span>6:30pm Sunday on<br><span class=\"visit\">Visit the Official Website</span></span>",
									"url": "http://agt.ninemsn.com.au"
								},
								"primaryColour": "#df0000",
								"seasons": null,
								"slug": "show/australiasgottalent",
								"title": "Australia's Got Talent",
								"tvChannel": "Channel 9"
							}
						],
						"skip": 0,
						"take": 1,
						"totalRecords": 1,
					}`,
			),
			want:    []model.StanleyResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan interface{})
			got := ConcurrentParser(done, tt.stream)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("ConcurrentParser() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Responses, tt.want) {
				t.Errorf("ConcurrentParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	tests := []struct {
		name    string
		jstr    string
		want    []model.StanleyResponse
		wantErr bool
	}{
		{
			"malformed json 3",
			`{
				"payload": [
				  {
					"Country": "UK",
					"Description": "What's life like when you have enough children to field your own football team?",
					"Drm": true,
					"EpisodeCount": 3,
					"Genre": "Reality",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "",
					  "Date": "",
					  "Html": "",
					  "Url": ""
					},
					"PrimaryColour": "#ff7800",
					"Seasons": [
					  {
						"Slug": "show/16kidsandcounting/season/1"
					  }
					],
					"Slug": "show/16kidsandcounting",
					"Title": "16 Kids and Counting Tv",
					"Channel": "GEM"
				  },
				  {
					"Country": "",
					"Description": "",
					"Drm": false,
					"EpisodeCount": 0,
					"Genre": "",
					"Image": {
					  "ShowImage": ""
					},
					"Language": "",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "",
					  "Date": "",
					  "Html": "",
					  "Url": ""
					},
					"PrimaryColour": "",
					"Seasons": [],
					"Slug": "show/seapatrol",
					"Title": "Sea Patrol Tv",
					"Channel": "Channel 9"
				  },
				  {
					"Country": "USA",
					"Description": "The Taste puts 16 culinary competitors in the kitchen, where four of the World's most notable culinary masters of the food world judges their creations based on a blind taste. Join judges Anthony Bourdain, Nigella Lawson, Ludovic Lefebvre and Brian Malarkey in this pressure-packed contest where a single spoonful can catapult a contender to the top or send them packing.",
					"Drm": true,
					"EpisodeCount": 2,
					"Genre": "Reality",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/TheTaste1280.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "http://catchup.ninemsn.com.au/img/player/logo_go.gif",
					  "Date": "",
					  "Html": "<br><span class=\"visit\">Visit the Official Website</span></span>",
					  "Url": ":http://go.ninemsn.com.au/"
					},
					"PrimaryColour": "#df0000",
					"Seasons": [
					  {
						"Slug": "show/thetaste/season/1"
					  }
					],
					"Slug": "show/thetaste",
					"Title": "The Taste (Le Goût) Tv",
					"Channel": "GEM"
				  },
				  {
					"Country": "UK",
					"Description": "The series follows the adventures of International Rescue, an organisation created to help those in grave danger using technically advanced equipment and machinery. The series focuses on the head of the organisation, ex-astronaut Jeff Tracy, and his five sons who piloted the \"Thunderbird\" machines.",
					"Drm": true,
					"EpisodeCount": 24,
					"Genre": "Action",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/Thunderbirds_1280.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "",
					  "Date": "",
					  "Html": "",
					  "Url": ""
					},
					"PrimaryColour": "#0084da",
					"Seasons": [
					  { "Slug": "show/thunderbirds/season/1" },
					  { "Slug": "show/thunderbirds/season/3" },
					  { "Slug": "show/thunderbirds/season/4" },
					  { "Slug": "show/thunderbirds/season/5" },
					  { "Slug": "show/thunderbirds/season/6" },
					  { "Slug": "show/thunderbirds/season/8" }
					],
					"Slug": "show/thunderbirds",
					"Title": "Thunderbirds Tv",
					"Channel": "Channel 9"
				  },
				  {
					"Country": "USA",
					"Description": "A sleepy little village, Crystal Cove boasts a long history of ghost sightings, poltergeists, demon possessions, phantoms and other paranormal occurrences. The renowned sleuthing team of Fred, Daphne, Velma, Shaggy and Scooby-Doo prove all of this simply isn't real, and along the way, uncover a larger, season-long mystery that will change everything.",
					"Drm": true,
					"EpisodeCount": 4,
					"Genre": "Kids",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/ScoobyDoo1280.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "",
					  "Date": "",
					  "Html": "",
					  "Url": ""
					},
					"PrimaryColour": "#1b9e00",
					"Seasons": [
					  {
						"Slug": "show/scoobydoomysteryincorporated/season/1"
					  }
					],
					"Slug": "show/scoobydoomysteryincorporated",
					"Title": "Scooby-Doo! Mystery Incorporated Tv",
					"Channel": "GO!"
				  },
				  {
					"Country": "USA",
					"Description": "Toy Hunter follows toy and collectibles expert and dealer Jordan Hembrough as he scours the U.S. for hidden treasures to sell to buyers around the world. In each episode, he travels from city to city, strategically manoeuvring around reluctant sellers, abating budgets, and avoiding unforeseen roadblocks.",
					"Drm": true,
					"EpisodeCount": 2,
					"Genre": "Reality",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/ToyHunter1280.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "",
					  "Date": "",
					  "Html": "",
					  "Url": ""
					},
					"PrimaryColour": "#0084da",
					"Seasons": [
					  {
						"Slug": "show/toyhunter/season/1"
					  }
					],
					"Slug": "show/toyhunter",
					"Title": "Toy Hunter Tv",
					"Channel": "GO!"
				  },
				  {
					"Country": "AUS",
					"Description": "A series of documentary specials featuring some of the world's most frightening moments, greatest daredevils and craziest weddings.",
					"Drm": true,
					"EpisodeCount": 1,
					"Genre": "Documentary",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/Worlds1280.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "",
					  "Date": "",
					  "Html": "",
					  "Url": ""
					},
					"PrimaryColour": "#ff7800",
					"Seasons": [
					  {
						"Slug": "show/worlds/season/1"
					  }
					],
					"Slug": "show/worlds",
					"Title": "World's... Tv",
					"Channel": "Channel 9"
				  },
				  {
					"Country": "USA",
					"Description": "Another year of bachelorhood brought many new adventures for roommates Walden Schmidt and Alan Harper. After his girlfriend turned down his marriage proposal, Walden was thrown back into the dating world in a serious way. The guys may have thought things were going to slow down once Jake got transferred to Japan, but they're about to be proven wrong when a niece of Alan's, who shares more than a few characteristics with her father, shows up at the beach house.",
					"Drm": true,
					"EpisodeCount": 0,
					"Genre": "Comedy",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/TwoandahHalfMen_V2.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "http://catchup.ninemsn.com.au/img/player/Ch9_new_logo.gif",
					  "Date": "",
					  "Html": "Next episode airs: <span> 10:00pm Monday on<br><span class=\"visit\">Visit the Official Website</span></span>",
					  "Url": "http://channelnine.ninemsn.com.au/twoandahalfmen/}",
					  "PrimaryColour": "#ff7800",
					  "Seasons": [],
					  "Slug": "show/twoandahalfmen",
					  "Title": "Two and a Half Men Tv",
					  "Channel": "Channel 9"
					}
				  },
				  {
					"Country": "USA",
					"Description": "Simmering with supernatural elements and featuring familiar and fan-favourite characters from the immensely popular drama The Vampire Diaries, it's The Originals. This sexy new series centres on the Original vampire family and the dangerous vampire/werewolf hybrid, Klaus, who returns to the magical melting pot that is the French Quarter of New Orleans, a town he helped build centuries ago.",
					"Drm": true,
					"EpisodeCount": 1,
					"Genre": "Action",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/TheOriginals1280.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "http://catchup.ninemsn.com.au/img/player/logo_go.gif",
					  "Date": "",
					  "Html": "<br><span class=\"visit\">Visit the Official Website</span></span>",
					  "Url": "http://go.ninemsn.com.au/"
					},
					"PrimaryColour": "#df0000",
					"Seasons": [
					  {
						"Slug": "show/theoriginals/season/1"
					  }
					],
					"Slug": "show/theoriginals",
					"Title": "The Originals Tv",
					"Channel": "GO!"
				  },
				  {
					"Country": "AUS",
					"Description": "Join the most dynamic TV judging panel Australia has ever seen as they uncover the next breed of superstars every Sunday night. UK comedy royalty Dawn French, international pop superstar Geri Halliwell, (in) famous Aussie straight-talking radio jock Kyle Sandilands, and chart -topping former AGT alumni Timomatic.",
					"Drm": false,
					"EpisodeCount": 0,
					"Genre": "Reality",
					"Image": {
					  "ShowImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/AGT.jpg"
					},
					"Language": "English",
					"NextEpisode": {
					  "Channel": "",
					  "ChannelLogo": "http://catchup.ninemsn.com.au/img/player/Ch9_new_logo.gif",
					  "Date": "",
					  "Html": "Next episode airs:<span>6:30pm Sunday on<br><span class=\"visit\">Visit the Official Website</span></span>",
					  "Url": "http://agt.ninemsn.com.au"
					},
					"PrimaryColour": "#df0000",
					"Seasons": [],
					"Slug": "show/australiasgottalent",
					"Title": "Australia's Got Talent Tv",
					"Channel": "Channel 9"
				  }
				]
			  }
			  `,
			[]model.StanleyResponse{
				{
					Image: "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg",
					Slug:  "show/16kidsandcounting",
					Title: "16 Kids and Counting Tv",
				},
				{
					Image: "http://catchup.ninemsn.com.au/img/jump-in/shows/TheTaste1280.jpg",
					Slug:  "show/thetaste",
					Title: "The Taste (Le Goût) Tv",
				},
				{
					Image: "http://catchup.ninemsn.com.au/img/jump-in/shows/Thunderbirds_1280.jpg",
					Slug:  "show/thunderbirds",
					Title: "Thunderbirds Tv",
				},
				{
					Image: "http://catchup.ninemsn.com.au/img/jump-in/shows/ScoobyDoo1280.jpg",
					Slug:  "show/scoobydoomysteryincorporated",
					Title: "Scooby-Doo! Mystery Incorporated Tv",
				},
				{
					Image: "http://catchup.ninemsn.com.au/img/jump-in/shows/ToyHunter1280.jpg",
					Slug:  "show/toyhunter",
					Title: "Toy Hunter Tv",
				},
				{
					Image: "http://catchup.ninemsn.com.au/img/jump-in/shows/Worlds1280.jpg",
					Slug:  "show/worlds",
					Title: "World's... Tv",
				},
				{
					Image: "http://catchup.ninemsn.com.au/img/jump-in/shows/TheOriginals1280.jpg",
					Slug:  "show/theoriginals",
					Title: "The Originals Tv",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		done := make(chan interface{})
		defer close(done)

		responses := ConcurrentParser(done, []byte(tt.jstr))

		if !compare(responses.Responses, tt.want) {
			t.Errorf("%s failed\ngot: %+v\nexpected: %+v\n", tt.name, responses.Responses, tt.want)
		}
	}
}

func compare(a, b []model.StanleyResponse) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_CheckDuplicate(t *testing.T) {
	tt := []struct {
		name    string
		jsonStr string
		wantErr bool
	}{
		{
			"duplicate keys case 1",
			`{
				"payload": [
					{
						"bamba": 2,
						"la": 5,
					},
					{
						"laa": 1,
						"laa": 2,
					},
					],
					}`,
			true,
		},
		{
			"duplicate keys case 2",
			`{
						"payload": [
							{
								"bamba": 2,
								"bamba": 5,
							},
							{
								"bamba": 1,
								"la": 2,
							},
						],
					}`,
			true,
		},
		{
			"duplicate keys case 3",
			`[
				{
					"bamba": 1,
					"la": 2,
				},
				{
					"bamba": 2,
					"bamba": 5,
				},
			]`,
			true,
		},
	}

	for _, testCase := range tt {
		jsonData := strings.NewReader(testCase.jsonStr)
		decoder := json.NewDecoder(jsonData)

		res := make(map[string]int)
		if err := testLevel(decoder, res, 0); err != nil {
			t.Errorf("%s failed %v", testCase.name, err)
		}
	}
}
