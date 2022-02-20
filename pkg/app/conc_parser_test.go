package app

import (
	"reflect"
	"testing"

	"github.com/darragh-downey/stanley/pkg/model"
)

func TestConcurrentParser(t *testing.T) {
	type args struct {
		done   chan interface{}
		stream []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []model.StanleyResponse
		wantErr bool
	}{
		{
			name:    "empty case",
			args:    args{},
			want:    []model.StanleyResponse{},
			wantErr: false,
		},
		{
			name: "single case",
			args: args{
				done: make(chan interface{}),
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
			},
			want:    []model.StanleyResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConcurrentParser(tt.args.done, tt.args.stream)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
