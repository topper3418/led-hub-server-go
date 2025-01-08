package device

import (
    "database/sql"
    "your-project/src/db"
)

type Device struct {
    ID        int
    Mac       string
    Name      sql.NullString
    Type      string
    CurrentIP sql.NullString
    Removed   bool
}

func Create(d Device) (int64, error) {
    conn := db.GetConnection()
    res, err := conn.Exec(`INSERT INTO devices (mac, name, type, current_ip, removed)
                           VALUES (?, ?, ?, ?, ?)`,
        d.Mac, d.Name, d.Type, d.CurrentIP, d.Removed)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}

func Get(id int) (*Device, error) {
    conn := db.GetConnection()
    var d Device
    err := conn.QueryRow(`SELECT id, mac, name, type, current_ip, removed
                          FROM devices WHERE id = ?`, id).
        Scan(&d.ID, &d.Mac, &d.Name, &d.Type, &d.CurrentIP, &d.Removed)
    if err != nil {
        return nil, err
    }
    return &d, nil
}

func Update(d Device) error {
    conn := db.GetConnection()
    _, err := conn.Exec(`UPDATE devices 
                         SET mac = ?, name = ?, type = ?, current_ip = ?, removed = ?
                         WHERE id = ?`,
        d.Mac, d.Name, d.Type, d.CurrentIP, d.Removed, d.ID)
    return err
}

func Delete(id int) error {
    conn := db.GetConnection()
    _, err := conn.Exec(`DELETE FROM devices WHERE id = ?`, id)
    return err
}

func GetAll() ([]Device, error) {
    conn := db.GetConnection()
    rows, err := conn.Query(`SELECT id, mac, name, type, current_ip, removed FROM devices`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var devices []Device
    for rows.Next() {
        var d Device
        err := rows.Scan(&d.ID, &d.Mac, &d.Name, &d.Type, &d.CurrentIP, &d.Removed)
        if err != nil {
            return nil, err
        }
        devices = append(devices, d)
    }
    return devices, rows.Err()
}

