import { get,post } from '@/utils/axios'

export function GetSellerApi() {
  return get('/seller')
}

export function GetPlistApi(param) {
  return post('/plist', param)
}
