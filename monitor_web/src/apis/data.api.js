import { get, post } from '@/utils/axios';

export function GetSellerApi() {
  return get('/seller');
}

export function GetPlistApi(param) {
  return post('/plist', param);
}

export function SubmitApi(param) {
  return post('/submit', param);
}

export function AuditList(param) {
  return post('/audit-list', param);
}

export function Audit(param) {
  return post('/audit', param);
}
