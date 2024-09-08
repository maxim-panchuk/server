package initdata

import (
	"reflect"
	"testing"
)

const (
	_parseTestInitData = "query_id=AAHdF6IQAAAAAN0XohDhrOrc&user=%7B%22id%22%3A279058397%2C%22first_name%22%3A%22Vladislav%22%2C%22last_name%22%3A%22Kibenko%22%2C%22username%22%3A%22vdkfrost%22%2C%22language_code%22%3A%22ru%22%2C%22is_premium%22%3Atrue%7D&auth_date=1662771648&hash=c501b71e775f74ce10e377dea85a7ea24ecd640b223ea86dfe453e0eaed2e2b2&start_param=abc"
	_myInitData        = "user=%7B%22id%22%3A715219007%2C%22first_name%22%3A%22%D0%9C%D0%B0%D0%BA%D1%81%D0%B8%D0%BC%22%2C%22last_name%22%3A%22%D0%9F%D0%B0%D0%BD%D1%87%D1%83%D0%BA%22%2C%22username%22%3A%22panchuk_maksim%22%2C%22language_code%22%3A%22ru%22%2C%22is_premium%22%3Atrue%2C%22allows_write_to_pm%22%3Atrue%7D&chat_instance=-844359077145115223&chat_type=sender&auth_date=1725828398&hash=6ae60841e019e295b937a9d737b2f1bbf5d810d8bc4d90be58993d5964c77e4f"
)

type testParse struct {
	initData    string
	expectedErr error
	expectedRes InitData
}

var testsParse = []testParse{
	{
		initData:    _myInitData + ";",
		expectedErr: ErrUnexpectedFormat,
	},
	{
		initData: _myInitData,
		expectedRes: InitData{
			User: User{
				AllowsWriteToPm: true,
				ID:              715219007,
				FirstName:       "Максим",
				LastName:        "Панчук",
				Username:        "panchuk_maksim",
				LanguageCode:    "ru",
				IsPremium:       true,
			},
			CanSendAfterRaw: 0,
			Hash:            "6ae60841e019e295b937a9d737b2f1bbf5d810d8bc4d90be58993d5964c77e4f",
		},
	},
}

func TestParse(t *testing.T) {
	for _, test := range testsParse {
		if data, err := Parse(test.initData); err != test.expectedErr {
			t.Errorf("expected error to be %q. \nReceived %q", test.expectedErr, err)
		} else if !reflect.DeepEqual(data, test.expectedRes) {
			t.Errorf("expected result to be %+v. \nReceived %+v", test.expectedRes, data)
		}
	}
}
