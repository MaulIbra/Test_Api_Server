package utils

const (
	CREATE_USER  = "INSERT INTO m_user(id_user,nik,username,tgl_lahir,id_pekerjaan,id_pendidikan) values (?,?,?,?,?,?)"
	SELECT_USERS = `SELECT 
	u.id_user,
    u.nik,
    u.username,
    u.tgl_lahir,
    p.id_pendidikan,
    p.label_pendidikan,
    u.id_pekerjaan,
    pk.label_pekerjaan,
    u.user_status,
    u.created_date,
    u.update_date
    FROM m_user as u INNER JOIN m_pendidikan as p
    ON u.id_pendidikan = p.id_pendidikan inner join
    m_pekerjaan as pk on u.id_pekerjaan = pk.id_pekerjaan where u.user_status = 1 LIMIT ?,?`
	SELECT_USER_BY_ID = `SELECT 
	u.id_user,
    u.nik,
    u.username,
    u.tgl_lahir,
    p.id_pendidikan,
    p.label_pendidikan,
    u.id_pekerjaan,
    pk.label_pekerjaan,
    u.user_status,
    u.created_date,
    u.update_date
    FROM m_user as u INNER JOIN m_pendidikan as p
    ON u.id_pendidikan = p.id_pendidikan inner join
    m_pekerjaan as pk on u.id_pekerjaan = pk.id_pekerjaan
	WHERE u.id_user = ?`
	UPDATE_USER = `UPDATE  m_user set nik=?,username=?,tgl_lahir=?,id_pekerjaan=?,id_pendidikan=? WHERE id_user = ?`
	DELETE_USER = `UPDATE m_user set user_status=? where id_user=?`
	SELECT_PEKERJAAN = `SELECT * FROM m_pekerjaan`
	SELECT_PENDIDIKAN = `SELECT * FROM m_pendidikan`

)
