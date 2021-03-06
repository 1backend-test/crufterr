import { Injectable } from '@angular/core';
import { NgClient } from '@1backend/ng-client';



export interface Test {
	name:	string;
	foods:	string[];
}


@Injectable()
export class Service {
  constructor(private ngClient: NgClient) {}

  GetHi(howManyTimes: number, allCool: boolean): Promise<string> {
    return this.ngClient.call<string>("crufterr", "test-service", "GET", "/hi", { "howManyTimes": howManyTimes, "allCool": allCool });
  }

  GetImportedHi(): Promise<void> {
    return this.ngClient.call<void>("crufterr", "test-service", "GET", "/imported-hi", {  });
  }

  GetSqlExample(): Promise<void> {
    return this.ngClient.call<void>("crufterr", "test-service", "GET", "/sql-example", {  });
  }

}
