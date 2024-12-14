package provider

func (p *Provider) FetchCount() (int, error) {
	var count int
	err := p.conn.QueryRow("SELECT value FROM counter LIMIT 1").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *Provider) IncreaseCount(value int) error {
	_, err := p.conn.Exec("UPDATE counter SET value = value + $1", value)
	return err
}
