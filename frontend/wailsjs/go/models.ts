export namespace main {
	
	export class Locale {
	    default: string;
	    supported: string[];
	
	    static createFrom(source: any = {}) {
	        return new Locale(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.default = source["default"];
	        this.supported = source["supported"];
	    }
	}
	export class Sponsor {
	    name: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new Sponsor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.url = source["url"];
	    }
	}
	export class ModMeta {
	    description: string;
	    author: string;
	    category: string[];
	    type: string;
	    version: string;
	    requires?: string[];
	    sponsor?: Sponsor;
	    locale?: Locale;
	
	    static createFrom(source: any = {}) {
	        return new ModMeta(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.author = source["author"];
	        this.category = source["category"];
	        this.type = source["type"];
	        this.version = source["version"];
	        this.requires = source["requires"];
	        this.sponsor = this.convertValues(source["sponsor"], Sponsor);
	        this.locale = this.convertValues(source["locale"], Locale);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ModData {
	    name: string;
	    repo: string;
	    downloads: number;
	    views: number;
	    meta: ModMeta;
	    locale: Locale;
	    featured: boolean;
	    _id: string;
	    version: number;
	    requires: string[];
	
	    static createFrom(source: any = {}) {
	        return new ModData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.repo = source["repo"];
	        this.downloads = source["downloads"];
	        this.views = source["views"];
	        this.meta = this.convertValues(source["meta"], ModMeta);
	        this.locale = this.convertValues(source["locale"], Locale);
	        this.featured = source["featured"];
	        this._id = source["_id"];
	        this.version = source["version"];
	        this.requires = source["requires"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

