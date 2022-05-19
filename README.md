# api



## Database

- `tables`

```sql
CREATE TABLE IF NOT EXISTS groups (
	id varchar primary key,
	name varchar not null,
	user_ids varchar -- user ids separated by comma.
);

CREATE TABLE IF NOT EXISTS users (
	id varchar primary key,
	username varchar not null,
	password varchar not null,
	email varchar
);
```
