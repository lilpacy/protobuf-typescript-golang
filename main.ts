import { Status } from './common/common_pb';

function doSomething(status: keyof typeof Status) {
  switch (status) {
    case 'ACTIVE':
      console.log('ステータスはアクティブです');
      break;
    case 'INACTIVE':
      console.log('ステータスは非アクティブです');
      break;
    case 'UNKNOWN':
      console.log('ステータスは不明です');
      break;
    default:
      const _: never = status;
      throw new Error('未知のステータス');
  }
}

doSomething('ACTIVE');
