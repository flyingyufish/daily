package redis

import (
	"errors"
	"fish/conn"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

const (
	ErrWriteHash = "[hash:%s]:write hash error,maybe:%v"
	ErrWirteList = "write list error"
	ErrWirteStr  = "write string error,maybe:%v"
	ErrWirteZset = "write zset error,maybe:%v"
	InfoDataNil  = "no data"
)

// EXPIRE 设置过期时间
func EXPIRE(key string, seconds interface{}) (bool, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("EXPIRE", redis.Args{}.Add(key).Add(seconds)...))

}

// DEL 删除某个键
func DEL(key string) (int64, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("DEL", key))
}

// EXISTS key是否存在
func EXISTS(key string) (exists bool, err error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("EXISTS", key))
}

// HSET 给hash添加字段
func HSET(key string, feild string, value interface{}) (int64, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("HSET", redis.Args{}.Add(key).AddFlat(feild).AddFlat(value)...))

}

// HEXISTS 判断hash某字段是否存在(HASH)
func HEXISTS(key, feild string) (bool, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("HEXISTS", redis.Args{}.Add(key).Add(feild)...))
}

// SISMEMBER 判断set里面是否存在这个元素(SET)
func SISMEMBER(key string, value interface{}) (bool, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("SISMEMBER", key, value))
}

// HMSET redis HMSET操作(HASH)
func HMSET(key string, obj interface{}) error {
	rc := conn.Get()
	defer rc.Close()
	result, err := redis.String(rc.Do("HMSET", redis.Args{}.Add(key).AddFlat(obj)...))
	if err != nil || result != "OK" {
		return fmt.Errorf(ErrWriteHash, key, err)
	}
	return nil
}

// HGET redisHGET操作(HASH)
func HGET(key string, obj interface{}) (int64, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("HGET", redis.Args{}.Add(key).AddFlat(obj)...))
}

// HINCRBY 给hash里面的某字段增加值，只能是整数
func HINCRBY(key, field string, obj interface{}) (int64, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("HINCRBY", redis.Args{}.Add(key).AddFlat(field).AddFlat(obj)...))
}

// HINCRBYFLOAT  给某个字段增加值,针对float(HASH)
func HINCRBYFLOAT(key, field string, obj interface{}) (int64, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("HINCRBYFLOAT", redis.Args{}.Add(key).AddFlat(field).AddFlat(obj)...))
}

// LPUSH 消息队列的写入操作
// key string 要写入的消息队列
// value interface{}写入的值
func LPUSH(key string, value interface{}) error {
	rc := conn.Get()
	defer rc.Close()
	resultID, err := redis.Int64(rc.Do("LPUSH", redis.Args{}.Add(key).Add(value)...))
	if resultID <= 0 {
		return fmt.Errorf(ErrWirteList)
	}
	return err
}

// BRPOP 消息队列的取出数据操作(阻塞)
// key 操作的队列zp
// timeout 超时
func BRPOP(key string, timeout int64) (strings []string, err error) {
	rc := conn.Get()
	defer rc.Close()
	var (
		values []interface{}
	)
	values, err = redis.Values(rc.Do("BRPOP", redis.Args{}.Add(key).AddFlat(timeout)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &strings)
	if err != nil {
		return
	}
	if len(strings) <= 0 {
		err = fmt.Errorf(InfoDataNil)
		return
	}
	return
}

// ZADD 添加到ZSET
// key: string sorted set的key
// field: 字段名
// weight: 权重即score
func ZADD(key string, field interface{}, weight interface{}) (bool, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("ZADD", redis.Args{}.Add(key).Add(weight).Add(field)...))

}

// ZCOUNT zset 判断某zset里某得分区间的总数
func ZCOUNT(key string, min, max interface{}) (count int64, err error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("ZCOUNT", redis.Args{}.Add(key).AddFlat(min).AddFlat(max)...))
}

// ZREVRANGEBYSCORE 获取权重值指定区间的值
// key zsetkey
//max 最大值
//min 最小值
func ZREVRANGEBYSCORE(key string, max interface{}, min interface{}) (strings []string, err error) {
	rc := conn.Get()
	defer rc.Close()
	var (
		values []interface{}
	)
	values, err = redis.Values(rc.Do("ZREVRANGEBYSCORE", redis.Args{}.Add(key).Add(max).Add(min)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &strings)
	if err != nil {
		return
	}
	if len(strings) <= 0 {
		err = fmt.Errorf(InfoDataNil)
		return
	}
	return
}

// ZINCRBY ReportUrgeZsetRds,key,1
func ZINCRBY(key string, value interface{}, score interface{}) (bool, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("ZINCRBY", redis.Args{}.Add(key).AddFlat(score).AddFlat(value)...))
}

// ZREMRANGEBYSCORE 删除指定分数范围的成员 [zset]
func ZREMRANGEBYSCORE(key string, min, max interface{}) (bool, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("ZREMRANGEBYSCORE", redis.Args{}.Add(key).AddFlat(min).AddFlat(max)...))
}

// ZRANGE 获取zset指定分数区间的fields
func ZRANGE(key string, min, max interface{}) (slice []string, err error) {
	var (
		values []interface{}
	)
	rc := conn.Get()
	defer rc.Close()
	values, err = redis.Values(rc.Do("ZRANGE", redis.Args{}.Add(key).Add(min).Add(max)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = fmt.Errorf(InfoDataNil)
		return
	}
	return
}

// ZCARD 计算zset内的个数
func ZCARD(key string) (count int64, err error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("ZCARD", redis.Args{}.Add(key)...))
}

// ZUNIONSTORE zset的并集操作
func ZUNIONSTORE(unionkey string, numkey interface{}, onekey, twokey, threekey string) (count int64, err error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("ZUNIONSTORE", redis.Args{}.Add(unionkey).AddFlat(numkey).AddFlat(onekey).AddFlat(twokey).AddFlat(threekey)...))
}

// ZSCORE 判断Zset下某value在不在
// ok 是否存在 T存在，F不存在
// score 得分，如果value存在必定大于0，当前得分
func ZSCORE(key string, obj interface{}) (ok bool, score int64, err error) {
	rc := conn.Get()
	defer rc.Close()
	resInt, err := redis.Int64(rc.Do("ZSCORE", redis.Args{}.Add(key).Add(obj)...))
	if err != nil {
		return false, 0, err
	}
	if resInt == 0 {
		return false, resInt, fmt.Errorf(InfoDataNil)
	}
	return true, resInt, nil //存在
}

// SET 在redis中写入string
// 返回值：设置成功时返回OK；seconds无效时，返回错误
// key string redis key
// value string key对应的值
func SET(key string, value interface{}) (ok bool, err error) {
	var (
		result string
	)
	rc := conn.Get()
	defer rc.Close()
	result, err = redis.String(rc.Do("SET", redis.Args{}.Add(key).Add(value)...))
	if err != nil && result != "OK" {
		return false, fmt.Errorf(ErrWirteStr, err)
	}
	return true, err
}

// SETEX 在redis中写入string
// 返回值：设置成功时返回OK；seconds无效时，返回错误
// key string redis key
// ex int  key的存活时间  单位s
// value interface{} key对应的值
func SETEX(key string, ex int64, value interface{}) (bool, error) {
	rc := conn.Get()
	defer rc.Close()
	result, err := redis.String(rc.Do("SETEX", redis.Args{}.Add(key).Add(ex).Add(value)...))
	if err != nil && result != "OK" {
		return false, fmt.Errorf(ErrWirteStr, err)
	}
	return true, nil
}

// INCR string类型的自增
func INCR(key string) (int, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int(rc.Do("INCR", redis.Args{}.Add(key)...))
}

// GET 从redis的string中取出值
//  返回值 get不会出现错误，当取到了值返回值，当没有取到就返回null，在这里，我们认为的当为null的时候就返回错误
func GET(key string) (string, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.String(rc.Do("GET", key))
}

// GETINT 获取一个字符串
func GETINT(key string) (int64, error) {
	rc := conn.Get()
	defer rc.Close()
	return redis.Int64(rc.Do("GET", key))
}

// ZREVRANGEWITHSCORES zset获取到并有权重
func ZREVRANGEWITHSCORES(zetKey string, start, end interface{}) (slice []string, err error) {
	var (
		values []interface{}
	)
	rc := conn.Get()
	defer rc.Close()
	values, err = redis.Values(rc.Do("ZREVRANGE", redis.Args{}.Add(zetKey).Add(start).Add(end).Add("withscores")...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = fmt.Errorf(InfoDataNil)
		return
	}
	return
}

// HGETALL redis HGETALL操作
func HGETALL(key, obj interface{}) (err error) {
	var v []interface{}
	rc := conn.Get()
	defer rc.Close()
	v, err = redis.Values(rc.Do("HGETALL", key))
	if err != nil {
		return
	}
	if len(v) == 0 {
		return errors.New(InfoDataNil)
	}

	err = redis.ScanStruct(v, obj)
	return
}

// GetZsetFields 顺序获取分页区间的zset里面的Fields
// key 需要操作的zset
// offset 上区间
// limit 下区间
func GetZsetFields(key string, offset, limit int64) (slice []string, err error) {
	var (
		values []interface{}
	)
	rc := conn.Get()
	defer rc.Close()
	values, err = redis.Values(rc.Do("ZRANGE", redis.Args{}.Add(key).Add(offset).Add(limit)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = errors.New(InfoDataNil)
		return
	}
	return
}

// ZREVRANGE 倒序获取分页区间的zset里面的Fields
// key 需要操作的zset
// offset 上区间
// limit 下区间
func ZREVRANGE(key string, offset, limit interface{}) (slice []string, err error) {
	var (
		values []interface{}
	)
	rc := conn.Get()
	defer rc.Close()
	values, err = redis.Values(rc.Do("ZREVRANGE", redis.Args{}.Add(key).Add(offset).Add(limit)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = errors.New(InfoDataNil)
		return
	}
	return
}

// Keys 获取Keys
func Keys(key string) (slice []string, err error) {
	rc := conn.Get()
	var (
		values []interface{}
	)
	defer rc.Close()
	values, err = redis.Values(rc.Do("keys", key))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = errors.New(InfoDataNil)
		return
	}
	return
}

//SADD 集合添加元素
func SADD(key string, obj interface{}) (err error) {
	rc := conn.Get()
	defer rc.Close()
	_, err = redis.Int64(rc.Do("SADD", redis.Args{}.Add(key).AddFlat(obj)...))
	return
}

// ZADDS ZSET 添加
func ZADDS(key string, mp map[float64]interface{}) (err error) {
	rc := conn.Get()
	defer rc.Close()
	var result int64
	result, err = redis.Int64(rc.Do("ZADD", redis.Args{}.Add(key).AddFlat(mp)...))
	if err != nil {
		return err
	}
	if result == 0 {
		return fmt.Errorf(ErrWirteZset, err)
	}
	return
}
