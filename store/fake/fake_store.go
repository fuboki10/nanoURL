package fake

var fakeRedis map[string]string

type FakeClient struct {
	fakeRedis map[string]string
}

func New() (*FakeClient, error) {
	fakeRedis = make(map[string]string)

	clinet := &FakeClient{fakeRedis: fakeRedis}

	return clinet, nil
}

func (client *FakeClient) Set(shortUrl string, originalUrl string) {
	client.fakeRedis[shortUrl] = originalUrl
}
	
func (client *FakeClient) Get(shortUrl string) string {
	return client.fakeRedis[shortUrl]
}


