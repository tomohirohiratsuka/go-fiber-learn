admin_db = db.getSiblingDB("admin");
admin_db.auth('root', 'root');

// create app db
app_db = admin_db.getSiblingDB("app");
app_db.createUser(
    {
        user: 'user',
        pwd: 'password',
        roles: [{role: 'readWrite', db: 'app'}],
    },
);

// create test db
test_db = admin_db.getSiblingDB("app_test");
test_db.createUser(
    {
        user: 'user',
        pwd: 'password',
        roles: [{role: 'readWrite', db: 'app_test'}],
    }
);
