/**
 * Generated by orval v7.9.0 🍺
 * Do not edit manually.
 * Tomodachi API
 * API for managing hubs in the Tomodachi application.
 * OpenAPI spec version: 1.0.0
 */

export interface Hub {
  id: number;
  name: string;
  /** @nullable */
  description?: string | null;
  /** @nullable */
  tags?: string[] | null;
  /** @nullable */
  event_type?: string | null;
  /** @nullable */
  languages?: string[] | null;
  /** @nullable */
  user_limit?: number | null;
  /** @nullable */
  access_rule?: string | null;
  /** @nullable */
  start_at?: number | null;
  /** @nullable */
  end_at?: number | null;
  /** @nullable */
  is_active?: boolean | null;
  /** @nullable */
  is_open?: boolean | null;
  /** @nullable */
  created_at?: number | null;
  /** @nullable */
  server_id?: number | null;
  /** @nullable */
  host_id?: number | null;
  [key: string]: unknown;
 }
