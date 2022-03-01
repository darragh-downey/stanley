package model

import "testing"

func TestStanleyRequest_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Country       string
		Description   string
		Drm           bool
		EpisodeCount  int
		Genre         string
		Image         StanleyImage
		Language      string
		NextEpisode   StanleyEpisode
		PrimaryColour string
		Seasons       []Season
		Slug          string
		Title         string
		TvChannel     string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "passing case",
			fields: fields{
				Country:      "USA",
				Description:  "The Taste puts 16 culinary competitors in the kitchen, where four of the World's most notable culinary masters of the food world judges their creations based on a blind taste. Join judges Anthony Bourdain, Nigella Lawson, Ludovic Lefebvre and Brian Malarkey in this pressure-packed contest where a single spoonful can catapult a contender to the top or send them packing.",
				Drm:          true,
				EpisodeCount: 2,
				Genre:        "Reality",
				Image: StanleyImage{
					ShowImage: "http://catchup.ninemsn.com.au/img/jump-in/shows/TheTaste1280.jpg",
				},
				Language: "English",
				NextEpisode: StanleyEpisode{
					Channel:     "",
					ChannelLogo: "http://catchup.ninemsn.com.au/img/player/logo_go.gif",
					Date:        "",
					Html:        "<br><span class=\"visit\">Visit the Official Website</span></span>",
					Url:         "http://go.ninemsn.com.au/",
				},
				PrimaryColour: "#df0000",
				Seasons: []Season{
					{
						Slug: "show/thetaste/season/1",
					},
				},
				Slug:      "show/thetaste",
				Title:     "The Taste (Le Goût)",
				TvChannel: "GEM",
			},
			args: args{
				data: []byte(`{
					"country": " USA",
					"description": "The Taste puts 16 culinary competitors in the kitchen, where four of the World's most notable culinary masters of the food world judges their creations based on a blind taste. Join judges Anthony Bourdain, Nigella Lawson, Ludovic Lefebvre and Brian Malarkey in this pressure-packed contest where a single spoonful can catapult a contender to the top or send them packing.",
					"drm": true,
					"episodeCount": 2,
					"genre": "Reality",
					"image": {
						"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/TheTaste1280.jpg"
					},
					"language": "English",
					"nextEpisode": {
						"channel": null,
						"channelLogo": "http://catchup.ninemsn.com.au/img/player/logo_go.gif",
						"date": null,
						"html": "<br><span class=\"visit\">Visit the Official Website</span></span>",
						"url": "http://go.ninemsn.com.au/"
					},
					"primaryColour": "#df0000",
					"seasons": [
						{
							"slug": "show/thetaste/season/1"
						}
					],
					"slug": "show/thetaste",
					"title": "The Taste (Le Goût)",
					"tvChannel": "GEM"
				},`),
			},
			wantErr: false,
		},
		{
			name:   "duplicate key case",
			fields: fields{},
			args: args{
				data: []byte(`{
					"country": " USA",
					"description": "The Taste puts 16 culinary competitors in the kitchen, where four of the World's most notable culinary masters of the food world judges their creations based on a blind taste. Join judges Anthony Bourdain, Nigella Lawson, Ludovic Lefebvre and Brian Malarkey in this pressure-packed contest where a single spoonful can catapult a contender to the top or send them packing.",
					"drm": true,
					"episodeCount": 2,
					"genre": "Reality",
					"image": {
						"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/TheTaste1280.jpg"
					},
					"language": "",
					"nextEpisode": {
						"channel": null,
						"channelLogo": "http://catchup.ninemsn.com.au/img/player/logo_go.gif",
						"date": null,
						"html": "<br><span class=\"visit\">Visit the Official Website</span></span>",
						"url": "http://go.ninemsn.com.au/"
					},
					"primaryColour": "#df0000",
					"seasons": [
						{
							"slug": "show/thetaste/season/1"
						}
					],
					"language": "English",
					"slug": "show/thetaste",
					"title": "The Taste (Le Goût)",
					"tvChannel": "GEM"
				},`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StanleyRequest{
				Country:       tt.fields.Country,
				Description:   tt.fields.Description,
				Drm:           tt.fields.Drm,
				EpisodeCount:  tt.fields.EpisodeCount,
				Genre:         tt.fields.Genre,
				Image:         tt.fields.Image,
				Language:      tt.fields.Language,
				NextEpisode:   tt.fields.NextEpisode,
				PrimaryColour: tt.fields.PrimaryColour,
				Seasons:       tt.fields.Seasons,
				Slug:          tt.fields.Slug,
				Title:         tt.fields.Title,
				TvChannel:     tt.fields.TvChannel,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("StanleyRequest.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
